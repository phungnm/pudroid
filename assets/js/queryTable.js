function readExcelFile(e) {
    var files = e.target.files, f = files[0];
    reader = new FileReader();
    
    reader.onload = function(e) {
      var data = new Uint8Array(e.target.result);
      var workbook = XLSX.read(data, {type: 'array'});
      var result = [];
      workbook.SheetNames.forEach(function(name) {
          result.push(XLSX.utils.sheet_to_json(workbook.Sheets[name],{header: 1}));
      });
	  console.log(workbook,result);
	  ip_rows = result[0];
	  if(ip_rows)
	  {
        $("#tool-check-ip-waiting").hide();

		  console.log(ip_rows);
	  }
    };

    reader.readAsArrayBuffer(f);
}
document.getElementById('tool-checkIP-input-file').addEventListener('change', function (e){
	if(e.target.files.length!=0){
        toastr.remove();
        $("#tool-check-ip-waiting").show();
        readExcelFile(e);
	}
}, false);