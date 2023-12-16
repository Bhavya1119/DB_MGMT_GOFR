Hello there, Bhavya Joshi this side !! 
I created this gofr application that allows you to insert and view the database entries for customers by creating endpoints. 

#1 Configure your database 
-change the db password and username for mysql server
-Also change the mysql port according to your installed port.

#2 Inserting entries 
- You can insert data using the curl queries :
  curl --location --request POST 'http://localhost:9000/customer/{id,name,email,phoneno}'
- You can also insert manually from mysql workbench or cli
- The data is also saved inside a csv file named "customers.csv"

#3 For displaying the entire table customers
  - You can use the endpoint "localhost:9000/customer"
  - Additionally i have provided a customer.html file that fetches the data from the csv file and allows user to view the tables with GUI

#4 Search by name
 - I have created an endpoint named search "localhost:9000/customer/search/{name}" with the help of which you can search the entire entry for that particular customer.


Created By Bhavya Joshi 
JIIT , NOIDA
20103191
