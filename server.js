//Dependencies
var koa = require('koa');
var logger = require('koa-logger');
var views = require("koa-views");
var Router = require("koa-router");
var serve = require("koa-static-folder");
var handlebars = require("koa-handlebars");
var marked = require('marked');
var parse = require('co-body');
var morph = require('morph');
var route = new Router();

//Environment file
var env = require('node-env-file'); //process.env.whatever
env(__dirname + '/.env');

//Database
var mongojs = require('mongojs');
var db = mongojs(process.env.HOST + '/' + process.env.NAME, ['posts']);

//Check connection
db.runCommand({ping:1}, function(err, res) {
    if(!err && res.ok) console.log("Connected to DB");
});

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
    var list;
    db.posts.find({}, function(err, docs){
        return docs;
    });

    yield this.render("index", {
        posts : list
    });
}

function *post(){
    var new_post = db.posts.findOne({link: this.params.link}, function(err, doc){
        return doc;
    });

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
    var post = yield parse(this);
    post.date = new Date;
    post.link = morph.toSnake(post.title);
    db.posts.insert(post);
    this.redirect('/');
}

//Set the port
app.listen(3000);
console.log("Listening on port 3000");