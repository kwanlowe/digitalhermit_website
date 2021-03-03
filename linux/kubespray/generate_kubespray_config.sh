declare -a IPS=($(gcloud compute instances list|awk 'NR>1{printf "%s ", $$5}'))
CONFIG_FILE=inventory/mycluster/hosts.yaml2 python3 src/kubespray/contrib/inventory_builder/inventory.py $${IPS[@]}

