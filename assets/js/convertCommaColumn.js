function columnToComma(column,delimiter)
{
	let comma_array = column.split('\n');
	comma_array.forEach(function(element,index) {
 comma_array[index] = element.trim();
});
	let comma = comma_array.join(delimiter);
	return comma;
}
function commaToColumn(comma,delimiter)
{
	let column_array = comma.split(delimiter);
	 column_array.forEach(function(element,index) {
 column_array[index] = element.trim();
});
	let column = column_array.join('\n');
	return column;
}



$(document).ready(function ($) {
	$('input[type=radio][name=delimiter_to_comma]').change(function() {
		  if (this.value == 1) {
		  	$("#delimiter-to-comma").val(", ");
		  	$("#delimiter-to-comma").change();
		  	$("#delimiter-to-comma").prop("disabled",true);

		  }
		  if (this.value == 2) {
		  	$("#delimiter-to-comma").val(";");
		  	$("#delimiter-to-comma").change();
		  	$("#delimiter-to-comma").prop("disabled",true);

		  }
		  if (this.value == 0) {
		  	$("#delimiter-to-comma").prop("disabled",false);

		  }
	});
	$('input[type=radio][name=delimiter_to_column]').change(function() {
		  if (this.value == 1) {
		  	$("#delimiter-to-column").val(", ");
		  	$("#delimiter-to-column").change();
		  	$("#delimiter-to-column").prop("disabled",true);
		  }
		  if (this.value == 2) {
		  	$("#delimiter-to-column").val(";");
		  	$("#delimiter-to-column").change();
		  	$("#delimiter-to-column").prop("disabled",true);
		  }
		  if (this.value == 0) {
		  	$("#delimiter-to-column").prop("disabled",false);
		  }
	});
	$("#data-to-comma,#delimiter-to-comma").on("keyup change",function(){
		let data = $("#data-to-comma").val();
		let delimiter = $("#delimiter-to-comma").val();
		let result = columnToComma(data,delimiter);
		$("#result-to-comma").val(result);
	});

	$("#data-to-column,#delimiter-to-column").on("keyup change",function(){
		let data = $("#data-to-column").val();
		let delimiter = $("#delimiter-to-column").val();
		let result = commaToColumn(data,delimiter);
		$("#result-to-column").val(result);
	});


	$("#copy-result-to-comma").on("click",function(){
		$("#copy-result-to-comma").html(`<i class="far fa-clipboard"></i> Copied!!!`);
			$("#copy-result-to-comma").prop("disabled", true);
			copyText($("#result-to-comma").val());
			setTimeout(function() {
				$("#copy-result-to-comma").prop("disabled", false);
				$("#copy-result-to-comma").html(`<i class="far fa-clipboard" ></i> Copy`);
			}, 1500);
	})


	$("#copy-result-to-column").on("click",function(){
	$("#copy-result-to-column").html(`<i class="far fa-clipboard"></i> Copied!!!`);
		$("#copy-result-to-column").prop("disabled", true);
		copyText($("#result-to-column").val());
		setTimeout(function() {
			$("#copy-result-to-column").prop("disabled", false);
			$("#copy-result-to-column").html(`<i class="far fa-clipboard"></i> Copy`);
		}, 1500);
	})
		



});