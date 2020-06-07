
(function($) {
    $("#zakaat_in_btc").hide();
    $(".rest").hide();
    $(".zkcss").hide();
    $("#zk_note").hide();

    var form = $("#signup-form");
    form.steps({
        headerTag: "h3",
        bodyTag: "fieldset",
        transitionEffect: "fade",
        labels: {
            previous : 'Previous',
            next : 'Submit',
            finish : 'Submit',
            current : ''
        },
        titleTemplate : '<div class="title"><span class="title-text">#title#</span><span class="title-number">0#index#</span></div>',
        onFinished: function (event, currentIndex)
        {
            console.log('Sumited');
        }
    });
    $("li").removeClass("disabled");    
    
})(jQuery);


var typed3 = new Typed('#typed', {
        strings: ['<i><b> Zakat</b></i>', '<i><b> Hajj</b></i>', '<i><b> Qurbani</b></i>', '<i><b> Fitra</b></i>'],
        typeSpeed: 100,
        backSpeed: 100,
        smartBackspace: true, // this is a default
        loop: true
      });

      function myFunction() {
      var x = document.getElementById("myTopnav");
      if (x.className === "topnav") {
        x.className += " responsive";
      } else {
        x.className = "topnav";
      }
    }
$("#btc").click(function(){
    $(".assetText").html("<img src='images/btc.png' width='20'>BTC<input type='hidden' class='asset' value='BTC'>");  
    console.log($(".asset").val());
})
$("#eth").click(function(){
    $(".assetText").html("<img src='images/eth.png' width='20'>ETH<input type='hidden' class='asset' value='ETH'>");  
    console.log($(".asset").val());
})

$("#usdt").click(function(){
    $(".assetText").html("<img src='images/usdt.png' width='20'>USDT<input type='hidden' class='asset' value='USDT'>");  
    console.log($(".asset").val());
})

$("#mkr").click(function(){
    $(".assetText").html("<img src='images/mkr.png' width='20'>MKR<input type='hidden' class='asset' value='MKR'>");  
    console.log($(".asset").val());
})

$("#bch").click(function(){
    $(".assetText").html("<img src='images/bch.png' width='20'>BCH<input type='hidden' class='asset' value='BCH'>");  
    console.log($(".asset").val());
})

$("#xrp").click(function(){
    $(".assetText").html("<img src='images/xrp.png' width='20'>XRP<input type='hidden' class='asset' value='XRP'>");  
    console.log($(".asset").val());
})

$("#hajj_btn").click(function(){
    var amt = $("#hajj_amount").val();
    var asset = $("#hajj_asset").val();

    console.log(asset);

    $.ajax({
        type: 'POST',
        url: "/hajj",
        data: {amt: amt, asset: asset},
    }).done(function(result){
        console.log(result);

         // var btc = Math.round(result * 100) / 100;
            var obj = JSON.parse(result);
            
            var total = Math.round(obj.hajj_in_USD* 100) / 100;
            
            var hajjNote = "<i>Note: Hajj is calculated, if an individual have $14000 USD or more because hajj is approx. $8000 USD maximum from anywhere around the world and remaining is $6000 USD for expenses. This calcualation may vary from place to place but taken as per the max amount required.</i>"
            console.log(total);
            if (total <= 0){
                $(".content").css("height", "365px");
                $("#hajj_result").html("Hajj is NOT compulsory");
                $("#hajj_note").html(hajjNote);
            } else {
                // $("#welcome_msg").hide();
                $(".content").css("height", "365px");
                // $(".zkcss").show();
                
                $("#hajj_result").html("Hajj is compulsory");
                $("#hajj_note").html(hajjNote);

            }
    })
})

// ajax call for zakaat
$("#zk_btn").click(function(){
    //button loader js


    var $this = $(this);
  $this.button('loading');
    setTimeout(function() {
       $this.button('reset');
   }, 1000);

    var amt = $("#amount").val();
    var asset = $(".asset").val();
    console.log(amt);
    console.log(asset);
    $.ajax({
        type: 'POST',
        url: '/zakaat',
        data: {amt:amt, asset: asset},
        
        }).done(function(result) { //use this
            console.log(result)
            // var btc = Math.round(result * 100) / 100;
            var obj = JSON.parse(result);
            
            var btc = Math.round(obj.zakaat_in_bitcoin* 100) / 100;
            
            var asset = obj.asset;

            var usd = Math.round(obj.zakaat_in_USD* 100) / 100;
            var gold = Math.round(obj.zakaat_in_gold* 100) / 100;
            var rm = Math.round(obj.zakaat_in_rm* 100) / 100;
            console.log(btc);
            if (btc < 0){
                $("#welcome_msg").hide();
                $(".zkcss").show();
                $("#zakaat_in_btc").show();
                $("#zakaat_in_btc").html(0.00+ " "+asset);
                $(".content").css("height", "365px");
                $(".rest").show();
                $(".rest").html("<span>"+0.00 +" in USD</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span>"+0.00 +"gm in GOLD</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span>"+0.00 +" in MYR</span>");
            } else {
                // $("#welcome_msg").hide();
                $(".content").css("height", "365px");
                $(".zkcss").show();
                $("#zakaat_in_btc").show();
                $("#zakaat_in_btc").html(btc+ " "+asset);

                $("#zk_note").show();
                $(".rest").show();
                $(".rest").html("<span>"+usd +" in USD</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span>"+gold +"gm in GOLD</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span>"+rm +" in MYR</span>");
            }

    })
})    


$(function() {
  $('a[href*=#]').on('click', function(e) {
    e.preventDefault();
    $('html, body').animate({ scrollTop: $($(this).attr('href')).offset().top}, 600, 'linear');
  });
});



