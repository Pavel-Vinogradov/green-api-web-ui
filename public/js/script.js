// Sync connection params to hidden inputs
document.getElementById('idInstance').addEventListener('input', function() {
    document.getElementById('connection-idInstance').value = this.value;
});

document.getElementById('apiTokenInstance').addEventListener('input', function() {
    document.getElementById('connection-apiTokenInstance').value = this.value;
});

// Sync message params
document.getElementById('chatId').addEventListener('input', function() {
    document.getElementById('message-chatId').value = this.value;
});

document.getElementById('message').addEventListener('input', function() {
    document.getElementById('message-message').value = this.value;
});

// Sync file params
document.getElementById('fileChatId').addEventListener('input', function() {
    document.getElementById('file-chatId-input').value = this.value;
});

document.getElementById('urlFile').addEventListener('input', function() {
    document.getElementById('file-urlFile').value = this.value;
});

document.getElementById('fileName').addEventListener('input', function() {
    document.getElementById('file-fileName').value = this.value;
});

document.getElementById('caption').addEventListener('input', function() {
    document.getElementById('file-caption').value = this.value;
});

// API call functions
async function callAPI(endpoint, data) {
    const idInstance = document.getElementById('idInstance').value;
    const apiTokenInstance = document.getElementById('apiTokenInstance').value;

    if (!idInstance || !apiTokenInstance) {
        alert('Please enter ID Instance and API Token Instance');
        return;
    }

    const loading = document.getElementById('loading');
    const response = document.getElementById('response');

    loading.classList.add('active');
    response.value = '';

    try {
        const response_data = {
            idInstance: idInstance,
            apiTokenInstance: apiTokenInstance,
            ...data
        };

        const res = await fetch(`/api/${endpoint}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(response_data)
        });

        const result = await res.json();
        response.value = JSON.stringify(result, null, 2);
    } catch (error) {
        response.value = `Error: ${error.message}`;
    } finally {
        loading.classList.remove('active');
    }
}

function callGetSettings() {
    callAPI('getSettings', {}).then(r => r);
}

function callGetStateInstance() {
    callAPI('getStateInstance', {}).then(r => r);
}

function callSendMessage() {
    const chatId = document.getElementById('chatId').value;
    const message = document.getElementById('message').value;

    if (!chatId || !message) {
        alert('Please enter Chat ID and Message');
        return;
    }

    callAPI('sendMessage', {
        chatId: chatId,
        message: message
    }).then(r => r);
}

function callSendFileByUrl() {
    const chatId = document.getElementById('fileChatId').value;
    const urlFile = document.getElementById('urlFile').value;
    const fileName = document.getElementById('fileName').value;
    const caption = document.getElementById('caption').value;

    if (!chatId || !urlFile || !fileName) {
        alert('Please enter Chat ID, File URL, and File Name');
        return;
    }

    callAPI('sendFileByUrl', {
        chatId: chatId,
        urlFile: urlFile,
        fileName: fileName,
        caption: caption
    }).then(r => r);
}
