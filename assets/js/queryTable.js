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
		  console.log(ip_rows);
	  }

  
      /* DO SOMETHING WITH workbook HERE */
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



jQuery('#file').change(function(){
    var file = document.getElementById('file').files[0];
    var progress = jQuery('#progress');

    if(file){
      var reader = new FileReader();
      var size = file.size;
      var chunk_size = Math.pow(2, 13);
      var chunks = [];

      var offset = 0;
      var bytes = 0;


      reader.onloadend = function(e){
        if (e.target.readyState == FileReader.DONE){
          var chunk = e.target.result;
          bytes += chunk.length;

          chunks.push(chunk);

          progress.html(chunks.length + ' chunks // ' + bytes + ' bytes...');

          if((offset < size)){
            offset += chunk_size;
            var blob = file.slice(offset, offset + chunk_size);

            reader.readAsText(blob);

          } else {
            progress.html("processing teh content...");

            var content = chunks.join("");

            alert("content is ready!");
            debugger;
          };
        }
      };

      var blob = file.slice(offset, offset + chunk_size);
      reader.readAsText(blob);
    }
  });