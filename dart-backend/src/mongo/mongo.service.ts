import { Db, MongoClient } from 'mongodb';
import process from 'node:process';

const MONGO_URL = process.env.mongodbUrl || 'mongodb://dart:d4rt@127.0.0.1:27017';
const MONGO_DB =  "dart";

const mongoDb = async (): Promise<Db> => {
    try {
        const client = new MongoClient(MONGO_URL);
        await client.connect();
        return client.db(MONGO_DB);
    } catch (err) {
        console.error('MongoDB connection error:', err);
        throw err;
    }
}

export default mongoDb;