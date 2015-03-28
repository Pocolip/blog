//Dependencies
var koa = require('koa');
var logger = require('koa-logger');
var views = require("koa-views");
var Router = require("koa-router");
var serve = require("koa-static-folder");
var handlebars = require("koa-handlebars");
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
        }

    }

}));

app.use(logger());

route.get("/", index);

function *index(){
    yield this.render("index", {title : "Home"});
}