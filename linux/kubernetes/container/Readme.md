
## Gcloud Registry Login

*  Configure GCloud project and credential
   * Go to the Project Page
   * Go to the IAM & Admin|Service Accounts page
   * Click `+Create Service Account`
   * Enter *Service Account Name* and *Description*
   * Clck Create
   * Under the *Service Account Permissions*, select *Cloud Build Service Account*
     This grants limited access to the registry.
   * Click *Continue*
   * Click *Done*
   * Under the *Service Accounts* page, click the menu under *Actions* and select *Create Key*
   * Select JSON for key type
   * Click *Create*
   * Save this key securely as it gives read/write access to your registry
     * For the purposes of this tutorial, the keyfile is assumed to be in ~/.ssh/gcloud-keyfile.json

* Return to the DigitalHermit github directory under the linux/kubernetes/container directory

* Test the login using the JSON key 
  `make gcloud-login CREDFILE=~/.ssh/gcloud-keyfile.json`
  
* Next, we configure docker/podman to use these credentials:
  `gcloud auth configure-docker`

* For the next step, we will need the GCloud project iD. This can be retrieved from the Project Home page in the browser or
  if you have an account with appropriate permission, frmo the gcloud command line utility:
    [klowe@elysium] gcloud projects list 
    PROJECT_ID                NAME               PROJECT_NUMBER
    anthos-playground-286117  anthos-playground  1034691625917

* Tag the image we created in the previous step:
  ` podman tag podman-webserver <GCloud Host>/<Project ID>/<Image Name>`
  For example:
  ` podman tag podman-webserver gcr.io/anthos-playground-286117/podman-webserver`
  

* Now we can push the image:  
  ` podman push gcr.io/anthos-playground-286117/podman-webserver`

  If all goes well, you should see the image pushed:

    Getting image source signatures
    Copying blob 843c3701e622 done
    Copying blob d0f104dc0a1f done
    Copying config d18db62e9e done
    Writing manifest to image destination
    Storing signatures

