//Dependencies
var koa = require('koa');
var logger = require('koa-logger');
var views = require("koa-views");
var Router = require("koa-router");
var serve = require("koa-static-folder");
var handlebars = require("koa-handlebars");
var serve = require('koa-static-folder');
var marked = require('marked');
var parse = require('co-body');
var morph = require('morph');
var route = new Router();

//External Data
var db = require("./db");


//Start the app
var app = koa();

//Set templating engine
app.use(handlebars({
    defaultLayout: "main",
    partialsDir : "./views",
    helpers:
    {
        sample: function(obj) {
            return obj.substring(0, 100) + "...";
        },
        debug: function(obj){
            console.log("Working.");
            console.log(obj);
        },
        add: function(obj){
            return obj;
        }

    }

}));

app.use(logger());
app.use(serve("./js"));
app.use(serve("./css"));


route.get("/", index);
route.get("/post/:link", post);
route.get("/new", add);
route.post("/create", create);
app.use(route.routes());


function *index(){
    yield this.render("index", {
        posts : db.posts
    });
}

function *post(){
    var new_post;
    for(var i = 0; i < db.posts.length; i++) {
        if (this.params.link == db.posts[i].link) {
            new_post = db.posts[i];
        }
    }
    yield this.render("post",{
        title: new_post.title,
        date: new_post.date,
        text: marked(new_post.text)
    });
}
function *add(){
    yield this.render("new");
}

function *create(){
    console.log("hi");

    var post = yield parse(this);
    db.posts.push(post);
    post.date = new Date;
    post.link = morph.toSnake(post.title);
    this.redirect('/');
}

//Set the port
app.listen(3000);
console.log("Listening on port 3000");