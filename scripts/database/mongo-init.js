print("\n ############ Init script started ############ \n")
const database = "rest-template"
const collectionName = "user"
var uuid = UUID()
    .toString('hex')
    .replace(/^(.{8})(.{4})(.{4})(.{4})(.{12})$/, '$1-$2-$3-$4-$5')

//user: admin
//password: admin
const user = {
    _id: uuid,
    username: "admin",
    password: "$2a$12$FEcwl6m6XDfKM9grMoaVTOi0a45oRf1/FJNzzYeQhreLM3oKXL11G",
    createdAt: new Date(),
    updatedAt: new Date(),
}

// Switch to the target database
db = db.getSiblingDB(database);

// Check if the 'user' collection exists
if (!db.getCollectionNames().includes(collectionName)) {
    print("\n ----- Creating database and user collection ----- \n");

    // Create the collection if it doesn't exist
    db.createCollection(collectionName);
    print("\n ----- User collection created ----- \n");

    // Create a unique index on the 'username' field
    db.user.createIndex({ "username": 1 }, { unique: true });
    print("\n ----- Username unique index created ----- \n");

    // Insert the first user
    db.user.insertOne(user);
    print("\n ----- First user inserted ----- \n");
} else {
    print("\n ----- Collection already exists, skipping initialization ----- \n");
}

print("\n ############ Init script finished ############ \n");