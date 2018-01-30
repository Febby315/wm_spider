"use strict";
phantom.outputEncoding = "utf-8";
//加载第三方库及配置
var webpage = require("webpage");
var system = require("system");
var config = require("../utils/config");

//获取传入参数并解析
var params = JSON.parse(system.args[1]||"{}");

//创建并配置抓取窗口
var page = webpage.create();
page.settings.loadImages = false; //禁止加载图片
page.customHeaders = { Referer: params.url };

//配置超时退出(放弃抓取并标记失败)
setTimeout(function() {
    json.status=99;
    send(json);
    phantom.exit(999);
}, 60*1000);

//初始化数据
var json = { status: 99, _id: params.id, version: params.version };

//数据传回服务器
function send(json) {
    var options = {
        operation: "POST", encoding: "utf8",
        headers: { "Content-Type": "application/json" },
        data: JSON.stringify(json)
    };
    //打开新窗口调用接口将数据传回服务器后退出当前进程
    var sendPage = webpage.create();
    sendPage.open(config.db_server_url + config.db_list_source_edit,options,function(status) {
        console.log(sendPage.content);
        setTimeout(phantom.exit(status.toUpperCase() === "SUCCESS"?0:1),1000);
    });
}
//发起页面并等待渲染后抓取页面
page.open(params.url, function(status) {
    if (status !== "success") {
        json.status=99;
        send(json);
    }else{
        //访问成功后等待15秒抓取页面
        setTimeout(function() {
            json["list_content"] = page.evaluate(function() { return document.getElementsByTagName('html')[0].outerHTML; });
            json.status = 1;
            send(json);
        }, 15000);
    }
});