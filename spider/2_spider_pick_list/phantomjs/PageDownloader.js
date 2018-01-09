"use strict";
var page = require('webpage').create();
var system = require('system');

var config = require('../utils/config');

phantom.outputEncoding="utf-8";
page.settings.loadImages = false;	//禁止加载图片

var params={};
function init(){
	 params=JSON.parse(system.args[1]);
}
init();
page.customHeaders = {
	"Referer":params.url,
};
		var json = {
				status:99,//失败
				_id:params.id,
				version:params.version
		};
setTimeout(function(){  //配置超时退出
	send(json);
	phantom.exit(999);
},3*60000);

function send(json){
	var settings = {
		  operation: "POST",
		  encoding: "utf8",
		  headers: {
			"Content-Type": "application/json"
		  },
		  data: JSON.stringify(json)
		};
		  console.log(json);
	page.open(config.db_server_url+config.db_list_source_edit,settings,function(status) {
		if(status.toUpperCase()==="SUCCESS"&&page.plainText.toUpperCase()==="SUCCESS"){
			//setTimeout(phantom.exit(0),1000);
		}else{
			//setTimeout(phantom.exit(1),1000);
		}
	});
}

////////////////////////////////////////////////////////////////////////////////
console.log('status',params.url);
//
page.open(params.url, function(status) {
	console.log('status',status);
  if (status !== 'success') {
	  console.log(json);
	phantom.exit();
  }else{
		try{
			setTimeout(function(){
				var ua = page.evaluate(function() {
					  return document.getElementsByTagName('html')[0].innerHTML;
					});
							json['list_content'] = ua;//爬取内容
							json['status'] = 1;//状态为已爬取
							send(json);
			},15000);
		}catch(e){
				send(json);
			console.log('phantomjs 出错了',e);
		}
		setTimeout(function(){
						phantom.exit(0);
		},35000);
  } 
  
});

//
