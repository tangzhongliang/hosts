$('.dropdown .dropdown-menu li').click(function() {
  		/* Act on the event */
  		accountType = $(this).find('a').text();
  		$(this).parent().parent().find('.dropdown-toggle').text(accountType)
  		// this.parent('.dropdown-toggle').val(accountType)
  	});