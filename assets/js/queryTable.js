var ip_rows = [];
var workbook;
function readExcelFile(e) {
    var files = e.target.files, f = files[0];
	reader = new FileReader();
    
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

            reader.readAsBinaryString(blob);

          } else {
            progress.html("processing teh content...");
            var content = chunks.join("");
             workbook = XLSX.read(content, {
                type: 'binary'
              });
         // works
         console.log(workbook['Sheets']['Sheet1']);
          };
        }
      };


      var blob = file.slice(offset, offset + chunk_size);
      reader.readAsBinaryString(blob);
    }
    
  });