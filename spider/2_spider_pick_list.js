const path = require('path');
const spawn = require("child_process").spawn;
const send = require("./utils/send");
const config = require("./utils/config");
// 获取需要处理的任务
var result = send.post(config.db_server_url + config.db_list_source_query, { conf: { status: 0 } });
var data = result.data;
var lsList=[];
// 遍历执行任务
for(let i=0,len=data.length;i<len;i++){
	try {
		let params = [
			path.join(__dirname,"./phantomjs/2_PageDownloader.js"),
			JSON.stringify({ id: data[i]._id, url: data[i].list_url.trim(), version: data[i].version })
		];
		let ls_temp = spawn("phantomjs", params);
		lsList.push({ id:data[i]._id,ls:ls_temp });
		console.log("OPEN[pid=%s]:%s",ls_temp.pid, data[i].list_url);
	}catch(err){
		console.error("ERROR:%s[%s]\n%s",data[i].list_url,data[i]._id,err);
	}
}
// 运行日志输出
lsList.forEach(function ({id:id,ls:ls},i) {
	ls.stdout.on("data", (msg) => {
		console.log("[%s]Phantom_Print:",ls.pid,String(msg));
	});
	ls.stderr.on("data", (err) => {
		console.error("[%s]Phantom_Error:",ls.pid,String(err));
	});
	ls.on("exit", (code,signal) => {
		console.log("[%s]Phantom_Exit:%s",ls.pid,code,String(signal));
	});
});