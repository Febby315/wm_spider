var fs = require("fs");
var send = require("./utils/send");
var config = require("./utils/config");
var uuid = require("node-uuid");
//var moment=require('moment');
var url = require("url");

var result = send.post(config.db_server_url + config.db_list_detail_cleaned_query,{ conf: { dealStatus: 0 } });
var data = result.data;
var date = new Date().format("YYYY-MM-DD");
for (let i = 0, len = data.length; i < len; i++) {
	var contentImageSour=[],listImageSour=[],listShowImage=[];
	contentImageSour = JSON.parse(data[i].contentImageSour||"[]");
	listImageSour = JSON.parse(data[i].listImageSour||"[]");
	attach_files_sour = JSON.parse(data[i].attach_files_sour||"[]");
	for (let j=0;j<contentImageSour.length;j++) {
		if (contentImageSour[j]&&contentImageSour[j].sour_url) {
			contentImageSour[j].sour_url=getAbsUrl(data[i].detail_url,contentImageSour[j].sour_url);
			downimg(contentImageSour[j],contentImageSour[j].sour_url);
		}
	}
	for (let j=0;j<listImageSour.length;j++) {
		if (listImageSour[j]&&listImageSour[j].sour_url) {
			listImageSour[j].sour_url=getAbsUrl(data[i].detail_url,listImageSour[j].sour_url);
			downimg(listImageSour[j], listImageSour[j].sour_url);
		}
	}
	for (let j=0;j<contentImageSour.length;j++) {
		if (contentImageSour[j]) { listShowImage.push({url: contentImageSour[j].url}); }
	}
	for (let j=0;j<listImageSour.length;j++) {
		if (listImageSour[j]) { listShowImage.push({ url: listImageSour[j].url }); }
	}
	for (let j=0;j<attach_files_sour.length;j++) {
		if (attach_files_sour[j]&&attach_files_sour[j].sour_url) {
			attach_files_sour[j].sour_url=getAbsUrl(data[i].detail_url,attach_files_sour[j].sour_url);
			downimg(attach_files_sour[j], attach_files_sour[j].sour_url);
		}
	}
	var param = {
		_id: data[i]._id,
		dealStatus: 1,
		version: data[i].version,
		contentImageSour: JSON.stringify(contentImageSour),
		listImageSour: JSON.stringify(listImageSour),
		listShowImage: JSON.stringify(listShowImage),
		attach_files_sour: JSON.stringify(attach_files_sour),
	};
	send.post(config.db_server_url + config.db_list_detail_cleaned_edit, param);
}
//获取绝对链接
function getAbsUrl(baseUrl,href) {
	if(!/^http/i.test(href)){
		if(/^\/\//.test(href)){
			href=url.parse(baseUrl).protocol + href
		}else{ href=url.resolve(baseUrl,href); }
	}
	return href;
}
//下载图片
function downimg(contentImageSourOne, src) {
	let result = send.post(config.db_server_down_url + config.downImg_findOne,{ conf: { img_src: src } });
	if (result && result.img_url) {
		contentImageSourOne.url = result.img_url;
	} else {
		let reg=/\.(jpg|png|gif|jpeg|bmp)/i.exec(src);
		let suffix=reg?reg[0]:".jpg";
		let filename = uuid.v1().replace(/-/g, "") + suffix;
		let filepath = date + "/" + new Date().format("YYYYMMDDHH") + "/";
		contentImageSourOne.url = "{IMG}/" + filepath + filename;
		let ret = send.post(config.db_server_down_url + config.download_do, { img_src: src, img_url: filepath, img_name: filename });
		console.log("下载状态", ret);
		send.post(config.db_server_down_url + config.downImg_add, { img_src: src, img_url: "{IMG}/" + filepath + filename, img_name: filename, dealStatus: ret });
		
	}
}
