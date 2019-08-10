package main
import (
   "github.com/aws/aws-sdk-go/aws"
   "github.com/aws/aws-sdk-go/aws/session"
   "github.com/aws/aws-sdk-go/aws/awserr"
   "github.com/aws/aws-sdk-go/service/secretsmanager"
   "fmt"
   "flag"
   "os"
)
func main() {
   var secretName,secretVersion,region string
   flag.StringVar(&secretName, "name", "", "Name of secret")
   flag.StringVar(&secretVersion, "version", "AWSCURRENT", "Version Stage (default: AWSCURRENT)")
   flag.StringVar(&region, "region", "us-east-1", "AWS Region (default: us-east-1)")
   flag.Parse()
   if len(secretName) == 0 {
       fmt.Println("Error: Secret name required.")
       os.Exit(1)
   }
   sn := secretName

   s, err := session.NewSessionWithOptions(session.Options{
          Config: aws.Config{
		    Region: aws.String(region),
	          },
          })

   sm := secretsmanager.New(s)
   output, err := sm.GetSecretValue(&secretsmanager.GetSecretValueInput{
        SecretId: &sn,
        VersionStage: aws.String(secretVersion),
        })

    if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
            case secretsmanager.ErrCodeDecryptionFailure:
                // Secrets Manager can't decrypt the protected secret text using the provided KMS key.
                fmt.Println(aerr.Error())
                fmt.Println("Secrets Manager could not decrypt the secret.")
            case secretsmanager.ErrCodeInternalServiceError:
                // An error occurred on the server side.
                fmt.Println( aerr.Error())
                fmt.Println("Server side error. Possible zombie apocalypse?")
            case secretsmanager.ErrCodeInvalidParameterException:
                // You provided an invalid value for a parameter.
                fmt.Println( aerr.Error())
                fmt.Println("Invalid parameter. Check inputs.")
            case secretsmanager.ErrCodeInvalidRequestException:
                // You provided a parameter value that is not valid for the current state of the resource.
                fmt.Println( aerr.Error())
            case secretsmanager.ErrCodeResourceNotFoundException:
            // We can't find the resource that you asked for.
                fmt.Println( aerr.Error())
                fmt.Println("Is your secret name correct?")
            }
        } else {
            // Print the error, cast err to awserr.Error to get the Code and
            // Message from an error.
            fmt.Println(err.Error())
        }
        os.Exit(1)
   }
   fmt.Println(*output.SecretString)
}
