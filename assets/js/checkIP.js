var checkIPTask;
var checkIPFile;
var ip_table="";
var ip_detail_table = "";
var ip_duplicate_table="";

var ip_rows =[];
var ip_duplicate_rows =[];

var ip_khong_bat = ["192.168.11.","103.78.78.","103.79.79.","103.199."];
var full_ip_khong_bat = ["171.244.170.210"];

var ip_dac_biet = ["103.199.32.","103.199.33.","10.65.","113.185."];
var full_ip_dac_biet = ["192.168.11.226","192.168.11.227","103.199.32.58","103.78.78.253","27"];

var ip_warning = ["113.185."];
var full_ip_warning = [];
function handleFile(e) {
    var files = e.target.files, f = files[0];
    var reader = new FileReader();
    reader.onload = function(e) {
      var data = new Uint8Array(e.target.result);
      var workbook = XLSX.read(data, {type: 'array'});
      var result = [];
      workbook.SheetNames.forEach(function(name) {
          result.push(XLSX.utils.sheet_to_json(workbook.Sheets[name],{header: 1}));
      });
      console.log(workbook,result);
  checkIP(result[0]);
      /* DO SOMETHING WITH workbook HERE */
    };
    reader.readAsArrayBuffer(f);
}
function checkIP(ip_rows)
{
    ip_rows.forEach(function(ip,index) {
    ip_rows[index][5]= checkLevelIp(ip[2]);
    });
    ip_duplicate_rows = checkDuplicate(ip_rows);
    console.log("done process");
    ip_table.clear();
						
    ip_rows.forEach(function(ip,index) {
        let dom = ip[2]?ip[2]:"NULL";
        
        if(ip[5]==2)								
            dom = `<span onClick="viewDetailIp('${ip[2]}')" class="pointer text-info">${dom}</span>`;
        if(ip[5]==1)								
            dom = `<span onClick="viewDetailIp('${ip[2]}')" class="pointer text-warning">${dom}</span>`;
        if(ip[5]==0)								
            dom = `<span onClick="viewDetailIp('${ip[2]}')" class="pointer text-danger">${dom}</span>`;
        ip_table.row.add(
            [
            ip[0]?ip[0]:"NULL",
            ip[1]?ip[1]:"NULL",
            dom,
            ip[3]?ip[3]:"NULL",
            ip[4]?ip[4]:"NULL"
            ])
    })
    ip_table.draw();
    console.log("done 1");
    ip_duplicate_table.clear();
    // ip_duplicate_table.rows.add(ip_duplicate_rows);
    ip_duplicate_rows.forEach(function(ip,index) {
        if(ip.users.length==1) return false;
            let dom = ip.ip;
            if(ip.level==1)								
            dom = `<span onClick="viewDetailIp('${ip.ip}')" class="pointer text-warning">${dom}</span>`;
            if(ip.level==0)								
            dom = `<span onClick="viewDetailIp('${ip.ip}')" class="pointer text-danger">${dom}</span>`;
        ip_duplicate_table.row.add(
            [
            dom,
            ip.users?ip.users.length:"NULL",
            `<span id="user-list-${index}">`+(ip.users?ip.users.join(', '):"NULL")+`</span>`+`<span target="#user-list-${index}" class="ml-1 btn-copy badge badge-danger pointer"><i class="far fa-clipboard"></i> Copy</span>`,
            `<button onClick="viewDetailIp('${ip.ip}')" class="btn btn-info btn-sm ">detail</button>`
            ]); 
    })
    ip_duplicate_table.draw();
    console.log("done 2");

    $("#tool-checkIP-submit").html("Load file");
    $("#tool-checkIP-submit").prop("disabled",false);
}
function checkLevelIp(ip)
{
	
	for(let i=0;i<ip_khong_bat.length;i++)
	{
			if(ip.indexOf(ip_khong_bat[i])==0) return 2;
	}
	for(let i=0;i<full_ip_khong_bat.length;i++)
	{
			if(ip==full_ip_khong_bat[i]) return 1;
	}
	for(let i=0;i<ip_dac_biet.length;i++)
	{
			if(ip.indexOf(ip_dac_biet[i])==0) return 2;
	}
	for(let i=0;i<full_ip_dac_biet.length;i++)
	{
			if(ip==full_ip_dac_biet[i]) return 2;
	}

	for(let i=0;i<ip_warning.length;i++)
	{
			if(ip.indexOf(ip_warning[i])==0) return 1;
	}
	return 0;
}
function viewDetailIp(target_ip)
{
	ip_detail_table.clear();
	$("#ip-detail-modal").modal("show");

	ip_rows.forEach(function(ip) {
		if(ip[2]==target_ip)
		{

			let dom = ip[2]?ip[2]:"NULL";
								
								if(ip[5]==2)								
									dom = `<span class="text-info">${dom}</span>`;
								if(ip[5]==1)								
									dom = `<span class="text-warning">${dom}</span>`;
								if(ip[5]==0)								
									dom = `<span class="text-danger">${dom}</span>`;
			ip_detail_table.row.add([
									ip[0]?ip[0]:"NULL",
									ip[1]?ip[1]:"NULL",
									dom,
									ip[3]?ip[3]:"NULL",
									ip[4]?ip[4]:"NULL"
									]);
		}
	});
	
	ip_detail_table.draw();	
}
function checkDuplicate(rows)
{
	var result=[];
	rows.forEach(function(row) {
		if(row[5]==2) return false;
		let item = []; //tao item moi
		item.level=row[5];
		item.ip = row[2];
		item.users = [row[0]];
		let check_exist = 0;
		for(let i=0;i<result.length;i++)
		{
			if(result[i].ip==item.ip)
			{
				let check_user_exist =0;
				result[i].users.forEach(function(user) {
					if(item.users[0]==user)
					{
						check_user_exist = 1;
					}
				});
				if(check_user_exist==0) 
					result[i].users.push(item.users[0]);
				check_exist = 1;
			}
		}

		if(check_exist==0) result.push(item);
	});

							
		return result;
}

document.getElementById('tool-checkIP-input-file').addEventListener('change', handleFile, false);

$(document).ready(function ($) {
	jQuery.fn.dataTable.Api.register('processing()', function(show) {

		return this.iterator('table', function(ctx) {
			ctx.oApi._fnProcessingDisplay(ctx, show);
		});
	});

var text_ip_2 = "";
ip_khong_bat.forEach(function(ip,index)
{
	text_ip_2= text_ip_2+ip+"xxx";
	if(index!=ip_khong_bat.length-1||full_ip_khong_bat.length!=0)
	text_ip_2+=" | ";
});
full_ip_khong_bat.forEach(function(ip,index)
{
	text_ip_2= text_ip_2+ip;
	if(index!=full_ip_khong_bat.length-1)
	text_ip_2+=" | ";
});
if(ip_dac_biet.length!=0&&full_ip_dac_biet.length!=0)
	text_ip_2+=" | ";
ip_dac_biet.forEach(function(ip,index)
{
	text_ip_2= text_ip_2+ip+"xxx";
	if(index!=ip_dac_biet.length-1||full_ip_dac_biet.length!=0)
	text_ip_2+=" | ";
});
full_ip_dac_biet.forEach(function(ip,index)
{
	text_ip_2= text_ip_2+ip;
	if(index!=full_ip_dac_biet.length-1)
	text_ip_2+=" | ";
});

var text_ip_1 = "";
ip_warning.forEach(function(ip,index)
{
	text_ip_1= text_ip_1+ip+"xxx";
	if(index!=ip_warning.length-1||full_ip_warning.length!=0)
	text_ip_1+=", ";
});
full_ip_warning.forEach(function(ip,index)
{
	text_ip_1= text_ip_1+ip;
	if(index!=full_ip_warning.length-1)
	text_ip_1+=", ";
});


$("#list-ip-2").html(text_ip_2);
$("#list-ip-1").html(text_ip_1);






    ip_detail_table = $('#ip-detail-table').DataTable( {
    	    	  "searchHighlight": true,
    	"language": {

			"emptyTable": "Load file with data to view",
			"processing": '<i class="fa-2x fas fa-spinner fa-spin"></i>',
		},
		"dom": 'Bfrtip',
	      buttons: [
	      
            {
                extend: 'csvHtml5',
                title: 'Data export'
            },
            {
                extend: 'excelHtml5',
                title: 'Data export'
            },
             'copy'
        ],
		"processing":true
    } );


    ip_table = $('#ip-table').DataTable( {
 "searchHighlight": true,
    	"language": {
			"emptyTable": "Load file with data to view",
			"processing": '<i class="fa-2x fas fa-spinner fa-spin"></i>',

		},
		"dom": 'Bfrtip',
	      buttons: [
	      
            {
                extend: 'csvHtml5',
                title: 'Data export'
            },
            {
                extend: 'excelHtml5',
                title: 'Data export'
            },
             'copy'
        ],
		"processing":true
    } );

   ip_duplicate_table = $('#ip-duplicate-table').removeAttr('width').DataTable( {
   		 "searchHighlight": true,
   		"columns": [
		    { "width": "20%"},
		    { "width": "15%" },
		    { "width": "50%" },
		    { "width": "15%", "className": "text-center" }
  		],
   		"language": {
			"emptyTable": "No duplicate IP",
			"processing": '<i class="fa-2x fas fa-spinner fa-spin"></i>'
		},
				"dom": 'Bfrtip',
	       buttons: [
            {
                extend: 'csvHtml5',
                title: 'Data export'
            },
            {
                extend: 'excelHtml5',
                title: 'Data export'
            },
             'copy'
        ],
  		  "fixedColumns": true,
		"processing":true,
		 "order": [[ 1, "desc" ]]
    } );
   	$("#ip-duplicate-table-link").on('shown.bs.tab', function (e) {
   		ip_duplicate_table.columns.adjust();
       });

})