#!/bin/sh
#usage precopy.sh <ContainerName> <RemoteLocation> <page-server port>
if [ "$#" -ne 3 ] ; then 
	echo "usage precopy.sh <ContainerName> <RemoteLocation> <page-server port>"
	exit 1
fi
#renaming & formatting variables
containerName=$1
remoteHost=$2
port=$3
containerImageName=$( docker ps -a --filter name=$containerName --format \"{{.Image}}\" )
username=${remoteHost%@*}
ip=${remoteHost#*@}
containerImageName="${containerImageName%\"}"
containerImageName="${containerImageName#\"}"
check=$( ssh -i /root/.ssh/id_rsa $remoteHost docker images --format \"{{.Repository}}\"  $containerImageName)

#checking if image exists on remote host

if [ "$check" != "$containerImageName" ]
then
	echo "image not present on remote host"
	exit 1;
fi

#making directories for saving

mkdir image
directory=$( ssh $remoteHost pwd )
ssh $remoteHost mkdir image

#pre-dump 20 times

for i in `seq 1 20`; 
do
	echo "pre-dump $i"	
	ssh $remoteHost nohup criu page-server --images-dir $directory/image --port $port &
	sleep 2
	docker checkpoint --leave-running --predump --track-mem --image-dir=$( pwd )/image/ --page-server --address=$ip --port=$port $containerName
done

#final dump

ssh $remoteHost nohup criu page-server --images-dir $directory/image --port $port &
sleep 2
docker checkpoint --image-dir=$( pwd )/image/ --prev-image-dir=$( pwd )/image/ --page-server --address=$ip --port=$port $containerName

#sending rest of the files to the server

scp $(pwd)/image/* $remoteHost:$directory/image/

#restarting container

ssh $remoteHost docker create --name=$containerName $containerImageName
ssh $remoteHost docker restore --force=true --image-dir=$directory/image $containerName 

#housekeeping and deleting the files

rm -rf $( pwd )/image
ssh $remoteHost rm -rf $directory/image
