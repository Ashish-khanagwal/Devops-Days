const mongoose = require("mongoose");

const postSchema = new mongoose.Schema({
  title: {
    type: String,
    required: [true, "Post must have a title"],
  },
  body: {
    type: String,
    required: [true, "Post must have a body"],
  },
});

const Post = mongoose.model("Port, postSchema");
module.exports = Post;