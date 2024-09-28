print("\n ############ Init script started ############ \n")
const database = "rest-template"
var uuid = UUID()
    .toString('hex')
    .replace(/^(.{8})(.{4})(.{4})(.{4})(.{12})$/, '$1-$2-$3-$4-$5')

//user: admin
//password: admin
const user = {
    _id: uuid,
    createAt: new Date(),
    updateAt: new Date(),
    username: "admin",
    password: "$2a$12$FEcwl6m6XDfKM9grMoaVTOi0a45oRf1/FJNzzYeQhreLM3oKXL11G",
}

db = db.getSiblingDB(database);
print("\n ----- Rest template database created ----- \n")

db.createCollection('user');
print("\n ----- User collection created ----- \n")

db.user.createIndex( { "username": 1 }, { unique: true } )
print("\n ----- Username unique index created ----- \n")

db.user.insertOne(user);
print("\n ----- First user inserted ----- \n")

print("\n ############ Init script finished ############ \n")