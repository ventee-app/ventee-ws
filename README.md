## Ventee backend

Backend for the Ventee project

### Browser connection

```javascript
function op() {
  const ws = new WebSocket('ws://localhost:9099');
  ws.onopen = (x) => {
    console.log('opened', x);
    ws.onclose = (r) => console.log('closed', r);

    ws.onmessage = (m) => {
      console.log('message', m);
      const parsed = JSON.parse(m.data);
      if (parsed.event === 'register-connection') {
        console.log('assigned ID:', parsed.data.connectionId);
      }
    };
    ws.onerror = (e) => console.log('error', e);

    ws.send('test');
  };
}
```