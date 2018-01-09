var send = require('./utils/send');
var config = require('./utils/config');
var spawn = require('child_process').spawn;
//获取urlList表内容的接口
var result=send.post(config.db_server_url+config.db_list_detail_source_query,{'conf':{'dealStatus':0}});
var data= result.data;
var lsList=[];
for(let i=0,len=data.length;i<len;i++){//遍历配置表获取过来的数据
	try{
		let ls_temp=spawn('phantomjs', [__dirname+'/phantomjs/PageDownloader.js',JSON.stringify({
			"id":data[i]._id,
			"url":data[i].detail_url,
			"version":data[i].version
		})]);
		lsList.push({id:data[i]._id,ls:ls_temp});
		console.log("OPEN in [pid=%s]:%s",ls_temp.pid, data[i].detail_url);
	}catch(e){
		console.error('ERROR['+data[i]._id+']:'+data[i].detail_url+'\n',e);
	}
}
lsList.forEach(function ({id:id,ls:ls},i) {
	ls.stdout.on('data', function (msg) {
		console.log('Phantom_Print['+ls.pid+']:',msg+"");
	});
	ls.stderr.on('data', function (err) {
		console.error('Phantom_Error['+ls.pid+']:',err+"");
	});
	ls.on("exit", (code,signal) => {
		console.log('Phantom_Exit['+ls.pid+']:',code,signal||"");
	});
});