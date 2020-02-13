var last_url="";
var total_url=0;
$(document).ready(function($) {
	
	$("#shorten-list").delegate('.shorten-copy', 'click', function(event) {
		let button =$(this);
		if(button.prop("disabled"));
		else
		{
					let ele =	button.attr("attr-target");

					var $temp = $("<input>");
$("body").append($temp);
$temp.val($("#"+ele).val()).select();
document.execCommand("copy");
$temp.remove();
	
button.prop("disabled",true);
button.html(`<i class="far fa-clipboard mr-3"></i> Copied`);
setTimeout(function(){
 button.html(`<i class="far fa-clipboard mr-3"></i> Copy `);
button.prop("disabled",false);

}, 1500);
		}


	});
	$("#shorten-submit").on('click',function(event){

		if($("#shorten-submit").prop("disabled"));
		else
		{
			$("#shorten-submit").prop("disabled",true)
				$.ajax({
					url: '/api/shortenUrl/add',
					async: true,
					type: 'post',
					data: `
					{
						"url": "${$("#url-input").val()}"
					}`,
					contentType: "application/json",
					success: function(result) {
						last_url = $("#url-input").val();
						show_url = last_url.substring(0,60);
						if(last_url.length>60)
						show_url+= "...";
						$("#shorten-list").append(` 
   <a target="_blank" title="${last_url}"  data-toggle="tooltip"  href="${last_url}"> ${show_url}</a>
			<div class="form-group">
              <div class="input-group  mb-4">
                <div class="input-group-prepend">
                  <span class="input-group-text"><i class="fas fa-link"></i></span>
                </div>
                <input id="shorten-input-${total_url}" value=" pudroid.cf/go/${result.ShortenUrl.code}" readonly class="shorten-input input-info  form-control" type="text">

                  <div class="input-group-append">
    <button  attr-target="shorten-input-${total_url}" class="shorten-copy btn btn-warning" type="button"><i class="far fa-clipboard mr-1"></i> Copy </button>
  </div>
              </div>
            </div>`)

						if(total_url==0) $("#shorten-result").show();
						total_url++;
						$.notify({
							message: 'Success',
						}, {
							type: 'success',
							newest_on_top: true,
						});

					},
						error: function(e) {
							$.notify({
							message: e.responseJSON.error,
						}, {
							type: 'danger',
							newest_on_top: true,
						});
					}
				});

		}

	});
	$("#url-input").keyup(function(){
		regex = RegExp(/(https?:\/\/)?([\w\-])+\.{1}([a-zA-Z]{2,63})([\/\w-]*)*\/?\??([^#\n\r]*)?#?([^\n\r]*)/g);
		let check = regex.test($(this).val());
		console.log(check,last_url!=$(this).val());
		if(last_url!=$(this).val()&&check) $("#shorten-submit").prop("disabled",false);
		else $("#shorten-submit").prop("disabled",true);
	});

});