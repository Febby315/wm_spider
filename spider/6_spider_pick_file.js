const url = require("url");
const uuid = require("uuid");
const moment = require('moment');
const send = require("./utils/send");
const config = require("./utils/config");
// 获取需要处理的任务
var result = send.post(config.db_server_url + config.db_list_detail_cleaned_query,{ conf: { dealStatus: 0 } });
var data = result.data;
var baseDir = moment().format("YYYY-MM-DD");
for (let i = 0, len = data.length; i < len; i++) {
	let listImageSour = JSON.parse(data[i].listImageSour||"[]");	//列表图片
	let contentImageSour = JSON.parse(data[i].contentImageSour||"[]");	//内容图片
	let attach_files_sour = JSON.parse(data[i].attach_files_sour||"[]");//附件
	//绝对化路径并下载
	for (let j=0;j<listImageSour.length;j++) {
		if (listImageSour[j]&&listImageSour[j].sour_url) {
			listImageSour[j].sour_url=url.resolve(data[i].detail_url,listImageSour[j].sour_url);
			downimg(listImageSour[j], listImageSour[j].sour_url);
		}
	}
	for (let j=0;j<contentImageSour.length;j++) {
		if (contentImageSour[j]&&contentImageSour[j].sour_url) {
			contentImageSour[j].sour_url=url.resolve(data[i].detail_url,contentImageSour[j].sour_url);
			downimg(contentImageSour[j],contentImageSour[j].sour_url);
		}
	}
	for (let j=0;j<attach_files_sour.length;j++) {
		if (attach_files_sour[j]&&attach_files_sour[j].sour_url) {
			attach_files_sour[j].sour_url=url.resolve(data[i].detail_url,attach_files_sour[j].sour_url);
			downimg(attach_files_sour[j], attach_files_sour[j].sour_url);
		}
	}
	//合并图片数组
	let listShowImage = listImageSour.map((v)=>v.url).concat(contentImageSour.map((v)=>v.url));
	//提交修改
	let param = {
		_id: data[i]._id,dealStatus: 1,version: data[i].version,
		contentImageSour: JSON.stringify(contentImageSour),
		listImageSour: JSON.stringify(listImageSour),
		listShowImage: JSON.stringify(listShowImage),
		attach_files_sour: JSON.stringify(attach_files_sour),
	};
	send.post(config.db_server_url + config.db_list_detail_cleaned_edit, param);
}
//下载图片
function downimg(fileObj, src) {
	let result = send.post(config.downImg_server + config.downImg_findOne,{ conf: { img_src: fileObj.sour_url } });
	if (result && result.img_url) {
		fileObj.url = result.img_url;
	} else {
		// let reg=/\.(jpg|png|gif|jpeg|bmp)/i.exec(fileObj.sour_url);
		// let suffix=reg?reg[0]:".jpg";
		let ns=url.parse(fileObj.sour_url).pathname.split(".");
		let suffix = ns>1?"."+ns.pop():"";
		let filename = uuid.v1().replace(/-/g, "") + suffix;
		let filepath = [baseDir,moment().format("YYYYMMDDHH")].join("/");
		fileObj.url = ["{IMG}",filepath,filename].join("/");
		let ret = send.post(config.downImg_server + config.downImg_down, { img_src: fileObj.sour_url, img_url: filepath, img_name: filename });
		console.log("下载状态", ret);
		send.post(config.downImg_server + config.downImg_add, { img_src: fileObj.sour_url, img_url: fileObj.url , img_name: filename, dealStatus: ret.state });
	}
}