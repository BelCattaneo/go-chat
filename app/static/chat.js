$( "#connect-btn" ).on( "click", function() {
    $( "#target" ).trigger( "click" );
    const userIdInput = $('#user_id');
    const roomIdInput = $('#room_id');


    console.log('user_id ', userIdInput.val())
    console.log('room_id ', roomIdInput.val())

    connectToWebSocket()
    
} );

function connectToWebSocket() {
    var webSocket = new WebSocket('ws://localhost:8000/upgrade?user_id=1&room_id=2');

    webSocket.onerror = function(event) {
        onError(event)
    };
    
    webSocket.onopen = function(event) {
        onOpen(event)
    };
    
    webSocket.onmessage = function(event) {
        onMessage(event)
    };
    
    function onMessage(event) {
        console.log("onMessage: ", event)
    }
    
    function onOpen(event) {
        console.log("onOpen: ", event)

    }
    
    function onError(event) {
        console.log("onError: ", event)
    }
}
