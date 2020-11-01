var splitFile;
$(document).ready(function ($) {
	

  $("#tool-split-form").submit(function (e) {
      e.preventDefault();
      let keepHeader = 0;
      if($("#keep-header-input").prop("checked")==true) keepHeader = 1;
      else  keepHeader =0;
                      if ($('#tool-split-input-file')[0].files.length !== 0) {
                    var form_data = new FormData();
                    splitFile =  $('#tool-split-input-file').prop('files')[0]['name'];
                    form_data.append('file', $('#tool-split-input-file').prop('files')[0]);
                    form_data.append('rows', $("#tool-split-input-rows").val());
                    form_data.append('keep_header',keepHeader);

        $("#tool-split-waiting").show();		
        $("#tool-split-result-el").hide();
      $.ajax({
            url:"/tools/splitExcel",
            type:"post",
            dataType: 'json',
            contentType: false,
            processData: false,
            data: form_data,
          success: function( result ) {
            if(result.status)
            {
              window.open(result.url, '_blank');
            }
          }
        });
    }
      
  });
})



