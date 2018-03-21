"use strict";
phantom.outputEncoding = "utf-8";
//加载第三方库及配置
var webpage = require("webpage");
var system = require("system");
var config = require("../utils/config");

//获取传入参数解析并初始化数据
var params = JSON.parse(system.args[1]||"{}");
var json = { dealStatus: 99, _id: params.id, version: params.version };

//配置超时退出(放弃抓取并标记失败)
setTimeout(function() {
    send(json);
}, 90*1000);

//数据传回服务器
function send(json) {
    var options = {
        operation: "POST", encoding: "utf8",
        headers: { "Content-Type": "application/json" },
        data: JSON.stringify(json)
    };
    //打开新窗口调用接口将数据传回服务器后退出当前进程
    var sendPage = webpage.create();
    sendPage.open(config.db_server_url + config.db_list_detail_source_edit,options,function(status) {
        setTimeout(phantom.exit(status=="success"?0:1),1000);
    });
}

//创建并配置抓取窗口
var page = webpage.create();
page.settings.loadImages = false; //禁止加载图片
page.customHeaders = { Referer: params.url };
//发起页面并等待渲染后抓取页面
page.open(params.url, function(status) {
    if (status !== "success") {
        send(json);
    }else{
        //访问成功后等待15秒抓取页面
        setTimeout(function() {
            json["content_sour"] = page.evaluate(function() { return document.getElementsByTagName('html')[0].outerHTML; });
            json.dealStatus = 1;
            send(json);
        }, 15000);
    }
});