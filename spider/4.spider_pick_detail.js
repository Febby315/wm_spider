const spawn = require("child_process").spawn;
const send = require("./utils/send");
const config = require("./utils/config");
//获取urlList表内容的接口
var result = send.post(config.db_server_url + config.db_list_detail_source_query,{ conf:{ dealStatus:0 } });
var data = result.data;
var lsList=[];
for(let i=0,len=data.length;i<len;i++){
	try {
		let params = [
			__dirname + "/phantomjs/4_PageDownloader.js",
			JSON.stringify({ id: data[i]._id, url: data[i].detail_url.trim(), version: data[i].version })
		];
		let ls_temp = spawn("phantomjs", params);
		lsList.push({ id:data[i]._id,ls:ls_temp });
		console.log("OPEN[pid=%s]:%s",ls_temp.pid, data[i].detail_url);
	}catch(err){
		console.error("ERROR:%s[%s]\n%s",data[i].detail_url,data[i]._id,err);
	}
}
lsList.forEach(function ({id:id,ls:ls},i) {
	ls.stdout.on("data", (msg) => {
		console.log("[%s]Phantom_Print:",ls.pid,msg||"[unknown msg]");
	});
	ls.stderr.on("data", (err) => {
		console.error("[%s]Phantom_Error:",ls.pid,err||"[unknown error]");
	});
	ls.on("exit", (code,signal) => {
		console.log("[%s]Phantom_Exit:%s",ls.pid,code,signal||"");
	});
});