"use strict";

function addEntry(s, isError) {
    let p = document.createElement('p');
    p.textContent = s;
    if(isError)
        p.classList.add('error');
    document.getElementById('box').appendChild(p);
}

async function getChat() {
    let resp;
    try {
        resp = await fetch('/chat/');
    } catch(e) {
        addEntry(e.toString(), true);
        return;
    }

    if(!resp.ok) {
        addEntry(resp.status + ' ' + resp.statusText, true);
        return;
    }

    let body;
    try {
        body = await resp.text();
    } catch(e) {
        addEntry(e.toString(), true);
        return;
    }

    let l = body.split('\n');
    l.pop();

    for(let i = 0; i < l.length; i++) {
        let resp;
        try {
            resp = await fetch("/chat/" + l[i]);
        } catch(e) {
            addEntry(e.toString(), true);
            continue;
        }
        if(!resp.ok) {
            addEntry(resp.status + ' ' + resp.statusText, true);
            continue;
        }
        let body;
        try {
            body = await resp.text();
        } catch(e) {
            addEntry(e.toString(), true);
            continue;
        }

        addEntry(body);
    }
}

async function handleInput(e) {
    e.preventDefault();
    var input = document.getElementById("input");
    var inputform = document.getElementById("inputform");
    var value = input.value;

    let resp;
    try {
        resp = await fetch("/chat/", {
            method: 'POST',
            body: value,
        });
    } catch(e) {
        addEntry(e.toString(), true);
        return;
    }

    if(!resp.ok) {
        addEntry(resp.status + ' ' + resp.statusText, true);
        return;
    }

    location.reload();
}

window.onload = getChat;
document.getElementById("inputform").addEventListener("submit", handleInput);
