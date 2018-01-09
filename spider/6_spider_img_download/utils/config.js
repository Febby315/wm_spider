const config={
	db_server_url:"http://localhost:5000",
	db_server_down_url:"http://localhost:3001",
	downImg_add:"/spiderdb_img_download/downImg/add",
	downImg_findOne:"/spiderdb_img_download/downImg/findOne",
   	db_list_detail_cleaned_query:"/spiderdb/listdetailcleaned/list",//
	db_list_detail_cleaned_edit:"/spiderdb/listdetailcleaned/edit",//
	download_do:"/spiderdb_img_download/updownload/download",//

}

//
function clone(myObj){ 
	if(typeof(myObj) != 'object') return myObj; 
	if(myObj == null) return myObj; 
	var myNewObj = new Object(); 
	for(var i in myObj) 
	{
		if(i!=='_id')
		myNewObj[i] = clone(myObj[i]); 
	}
 return myNewObj; 
};
Date.prototype.format = function (fmt) { //author: meizz 
	var o = {
		"M+": this.getMonth() + 1, //月份
		"d+": this.getDate(), //日
		"D+": this.getDate(), //日
		"H+": this.getHours(), //24小时
		"h+": this.getHours()%12, //12小时
		"m+": this.getMinutes(), //分
		"s+": this.getSeconds(), //秒
		"S+": this.getMilliseconds(), //毫秒
		"q+": Math.floor((this.getMonth() + 3) / 3) //季度
	};
	if (/(Y+)/igm.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
	for (var k in o){
		if (new RegExp("(" + k + ")","gm").test(fmt)){
			fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
		}
	}
	return fmt;
};

module.exports = config;
module.exports.clone= clone;