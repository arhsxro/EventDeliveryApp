# EventDeliveryApp

Event Delivery App -- a system that receives events from multiple users from an HTTP
endpoint and delivers them to multiple destinations.

Durability:

	First of all this is my Kafka cluster configurations : 1 Broker, 1 replica 1 partition 1 topic 1 consumer.
	I decided to go with that 

	I configured the Producer.RequiredAcks to sarama.WaitForAll so the producer will wait for acknowledgment from all replicas
	in the Kafka cluster before considering the message as successfully written. This ensures that the message is
	durably stored in the Kafka cluster even if a broker fails.

	Also by using sarama.SyncProducer, each message published to Kafka will be acknowledged synchronously, ensuring durability.
	If the message cannot be acknowledged by all replicas, the producer will retry a certain number of times based on the config.
	Producer.Retry.Max value before giving up.
	
	To achieve better results i should have used something like 3 brokers and 3 replicas.With this configuration I could have good fault tolerance. 
	If one broker goes down, the other two replicas can still serve data, ensuring high availability.

	Data Redundancy: The 3 replicas provide data redundancy, meaning my data is stored on multiple brokers, reducing the risk of data loss due to broker failures.

	
Maintain order:

	To ensure maintaining order for events of the same user I used 1 partition.

	Kafka guarantees order within a partition, so by assigning the same partition to events with the same user ID,I can ensure that events
	from the same user are processed in the order they were received.
	
	If i had more partitions i could still maintain order by using the user ID as the partition key.
	Events from the same user will be assigned to the same partition, ensuring that they are processed in order within that partition.
	

At-least-once delivery:

	To achieve this I have implemented a retry mechanism to handle any failures and esnure that the message will be delivered.
	
	
Retry backoff and limit:

	To achieve retry backoff and limit I could implement a retry mechanism with a backoff algorithm. The backoff algorithm increases the wait time between retries,
	allowing for a controlled and incremental retry process.I have used retry mechanisms in many cases with increasing delay in every retry.


Delivery isolation:

	To achive this I Implemented separate functions for delivering events to each destination. Each function handles the delivery logic specific to that destination.
	This includes failures, delays etc.. The code for each destination can be independent, allowing for isolated delivery to each destination.



-------Execution-------

First option: Docker

	I dockerized this app to make it easier to run it
	The steps to run this app are these :
	1)docker  build .
	2)docker-compose up -d
	3)Now that our services are up and running we should set the topic config in kafka. So we run this command:
		docker-compose exec kafka kafka-topics.sh --create --topic events_topic --partitions 1 --replication-factor 1 --bootstrap-server localhost:9093
		
	Note that the topic name, the replicas and the partitions are based on my implemantation and can be changed.

	Now all we have to do is to open the postman and send a POST request to http://localhost:8080/api/singleEvent with this Json body : {
		"userID" : "456",
		"payload" : "First event"
	}

Second option: 

	Another way to run this loccally is to download kafka.Then open 3 seperate git bashes on the kafka file and run these commands:
	1)bin/zookeeper-server-start.sh config/zookeeper.properties  With this command we start the zookeeper server which is essential for kafka to run properly.
	2)bin/kafka-server-start.sh config/server.properties  With this command we start the kafka server
	3)bin/kafka-topics.sh --create --topic events_topic --partitions 1 --replication-factor 1 --bootstrap-server localhost:9092  With this commanmd we
		configure out kafka topic etc..
		
	Now that we are all set we can run our go app with this command -> go run main.go
	
	Now all we have to do is to open the postman and send a POST request to http://localhost:8080/api/singleEvent with this Json body : {
		"userID" : "456",
		"payload" : "First event"
	}



