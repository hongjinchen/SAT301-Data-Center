# User 

Use `bcrypt` is a good way to hash passwords, which can crypt the password in the database to improve security. The salt is a random string that is used to hash the password. In the database, the salt is stored with the password. Furthermore, salt is stored in the database with the passwords, so that the password can be hashed only once. 

The process with crypt:

A password "pass" and two salts: "s1", "s2".

- hash(s1 + pass) = 123

- hash(s2 + pass) = 456

So you would have two stored records in your DB:

- s1$123

- s2$456

reference: 
[1] Why is golang package bcrypt able to retrieve the salt after hashing the password?
https://stackoverflow.com/questions/64781341/why-is-golang-package-bcrypt-able-to-retrieve-the-salt-after-hashing-the-passwor