
// $(document).ready(function ($) {

    Dropzone.optionsForElement("#splitExcelForm") = {
        init: function() {
          this.on("addedfile", function(file) { alert("Added file."); });
        }
    };
// })
