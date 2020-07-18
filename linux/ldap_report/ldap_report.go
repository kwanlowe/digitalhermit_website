package main

import (
    "fmt"
    "log"
    "flag"
//    "crypto/tls"
    "strconv"
    toml "github.com/BurntSushi/toml"
    ldap "github.com/go-ldap/ldap"
    df "github.com/rocketlaunchr/dataframe-go"
)

type tomlConfig struct {
        Title string
        AuthConfig authConfig `toml:"auth"`
        LdapConfig ldapConfig `toml:"ldap"`
}

type authConfig struct {
       Username string
       Password string
       Filter string
}

type ldapConfig struct {
       Filter string
       Server string
       BaseDN string
       Port int
       SearchBase string
}


func GetConfig(configFile string) *tomlConfig {
    config := new(tomlConfig)
    if _, err := toml.DecodeFile(configFile, &config); err != nil {
         fmt.Println(err)
         return config
    }
    return config
}

const (
    queryLimit = 100000
)

func main() {

    // var tlsConf *tls.Config
    // tlsConf = &tls.Config{InsecureSkipVerify: true} // FIXME

    //  Read the configuration file for setup
    var configFile string

    flag.StringVar(&configFile, "config", "auth.toml", "configuration file")
    flag.Parse()
    conf := GetConfig(configFile)
    username := conf.AuthConfig.Username
    password := conf.AuthConfig.Password
    filter := conf.LdapConfig.Filter
    server := conf.LdapConfig.Server
    searchbase := conf.LdapConfig.SearchBase
    port := conf.LdapConfig.Port


    dialurl := server + ":" + strconv.Itoa(port)

    conn, err := ldap.Dial("tcp", dialurl)

    // conn, err := ldap.DialTLS("tcp", dialurl, tlsConf)
    if err != nil {
        fmt.Println("dialURL failure")
        log.Fatal(err)
    }
    defer conn.Close()

    _, err = conn.SimpleBind(&ldap.SimpleBindRequest{
        Username: username,
        Password: password,
    })

    if err != nil {
        fmt.Println("SimpleBind failure")
        log.Fatal(err)
    }

    var pageSize uint32 = 32
    searchBase := searchbase
    pagingControl := ldap.NewControlPaging(pageSize)
    attributes := []string{
        "cn", "dn", "uid", "objectClass", "mail", "sn", "telephoneNumber",
    }

    controls := []ldap.Control{pagingControl}

    s01 := df.NewSeriesString( "cn", nil)
    s02 := df.NewSeriesString( "dn", nil)
    s03 := df.NewSeriesString( "uid", nil)
    s04 := df.NewSeriesString( "objectClass", nil)
    s05 := df.NewSeriesString( "mail", nil)
    s06 := df.NewSeriesString( "telephoneNumber", nil)

    df := df.NewDataFrame(s01, s02, s03, s04, s05, s06)


    for {
        request := ldap.NewSearchRequest(searchBase, ldap.ScopeWholeSubtree,
                                         ldap.DerefAlways, 0, 0, false, filter,
                                         attributes, controls)
        response, err := conn.Search(request)

        response.PrettyPrint(5)

        if err != nil {
            log.Fatalf("Failed to execute search request: %s", err.Error())
        }


        for _, value := range response.Entries {
            t01 := value.GetAttributeValue("cn")
            t02 := value.GetAttributeValue("dn")
            t03 := value.GetAttributeValue("uid")
            t04 := value.GetAttributeValue("objectClass")
            t05 := value.GetAttributeValue("mail")
            t06 := value.GetAttributeValue("telephoneNumber")
            df.Append(nil, t01, t02, t03, t04, t05, t06)
        }

        // In order to prepare the next request, we check if the response
        // contains another ControlPaging object and a not-empty cookie and
        // copy that cookie into our pagingControl object:
        updatedControl := ldap.FindControl(response.Controls, ldap.ControlTypePaging)
        if ctrl, ok := updatedControl.(*ldap.ControlPaging); ctrl != nil && ok && len(ctrl.Cookie) != 0 {
            pagingControl.SetCookie(ctrl.Cookie)
            continue
        }
        // If no new paging information is available or the cookie is empty, we
        // are done with the pagination.
        break
    }

    // fmt.Println(df)
    fmt.Print(df.Table())
}
