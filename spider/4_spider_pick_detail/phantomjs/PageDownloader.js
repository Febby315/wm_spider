"use strict";
phantom.outputEncoding="utf-8";

var Browser = require('webpage');
var system = require('system');
var config = require('../utils/config');

var params={};
function initParams(){  //初始化参数
	params=JSON.parse(system.args[1]||"{}");
}
initParams();

var json = {};
json['_id'] = params.id;//状态为已爬取
json['version'] = params.version;//状态为已爬取
json['dealStatus'] = 99;//爬取状态 默认爬取失败99

setTimeout(function(){  //配置超时退出
	send(json);
	phantom.exit(999);
},3*60000);

function send(json){
	var settings = {
		operation: "POST",
		encoding: "utf8",
		headers: { "Content-Type": "application/json" },
		data: JSON.stringify(json)
	};
	var sendPage = Browser.create();
	sendPage.open(config.db_server_url+config.db_list_detail_source_edit,settings,function(status) {
		if(status.toUpperCase()==="SUCCESS"){   //&&page.plainText.toUpperCase()==="SUCCESS"
			phantom.exit(0);//正常退出
		}else{
			phantom.exit(1);//接口访问失败
		}
	});
}

var page=Browser.create();
page.settings.loadImages = false;	//禁止加载图片

page.open(params.url, function(status) {
	  if (status !== 'success') {
	phantom.exit();
  }else{
		
		setTimeout(function(){
		var ua = page.evaluate(function() {
			  return document.getElementsByTagName('html')[0].innerHTML;
			});

					json['content_sour'] = ua;//爬取内容
					json['dealStatus'] = 1;//状态为已爬取
					json['detail_url'] = page.url;//有跳转的改成跳转后的地址
					send(json);
					console.log("url "+params.url);	


},15000);

setTimeout(function(){
				phantom.exit(0)
},25000);
  } 
  

});