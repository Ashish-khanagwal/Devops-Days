const express = require("express");
const mongoose = require("mongoose");
const redis = require("redis");
const session = require("express-session");
let RedisStore = require("connect-redis").default;

const {
  MONGO_USER,
  MONGO_PASSWORD,
  MONGO_IP,
  MONGO_PORT,
  REDIS_URL,
  REDIS_PORT,
  SESSION_SECRET,
} = require("./config/config");

let redisClient = redis.createClient({
  host: REDIS_URL,
  port: REDIS_PORT,
});

const app = express();

const connectWithRetry = () => {
  mongoose
    .connect(
      `mongodb://${MONGO_USER}:${MONGO_PASSWORD}@${MONGO_IP}:${MONGO_PORT}/?authSource=admin`
    )
    .then(() => console.log("Successfully connected to DB"))
    .catch((e) => {
      console.log(e);
      setTimeout(connectWithRetry, 5000);
    });
};

connectWithRetry();

app.use(
  session({
    store: new RedisStore({
      client: redisClient,
    }),
    secret: SESSION_SECRET,
    cookie: {
      secure: false,
      resave: false,
      saveUninitialized: false,
      httpOnly: true,
      maxAge: 30000,
    },
  })
);

app.use(express.json());

app.get("/", (req, res) => {
  res.send("<h2>Hello World</h2>");
});

app.listen(3000, () => console.log("Listening on port 3000"));
