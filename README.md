## Ventee backend

Backend for the Ventee project

### Browser connection

```javascript
function emulate() {
  const ws = new WebSocket('ws://localhost:9099');
  ws.onopen = () => console.log('opened');
  ws.onclose = (r) => console.log('closed:', r);
  ws.onmessage = ({ data = '' }) => {
    if (!data) { return null };

    const parsed = JSON.parse(data);

    if (parsed.event === 'register-connection') {
      const payload = JSON.parse(parsed.data);
      console.log('assigned ID:', payload.connectionId);
      this.connectionId = payload.connectionId;
    }
    
    if (parsed.event === 'request-contacts') {
      if (parsed.target !== this.connectionId) {
        return null;
      }
      console.log('request contacts to', parsed.target);
      ws.send(JSON.stringify({
        data: JSON.stringify({ some: 'data' }),
        event: 'transfer-contacts',
        issuer: ws.connectionId,
        target: parsed.issuer,
      }));
    }

    if (parsed.event === 'transfer-contacts') {
      if (parsed.target !== this.connectionId) {
        return null;
      }
      const payload = JSON.parse(parsed.data);
      console.log('received contacts:', payload);

      // wrap things up
      ws.send(JSON.stringify({
        event: 'transfer-complete',
        issuer: ws.connectionId,
        target: parsed.issuer,
      }));
    }

    if (parsed.event === 'transfer-complete') {
      if (parsed.target !== this.connectionId) {
        return null;
      }
      console.log('data transfer is complete');
    }
  };
  ws.onerror = (e) => console.log('error', e);

  setTimeout(() => ws.send(JSON.stringify({
    event: 'request-contacts',
    issuer: this.connectionId,
    target: this.connectionId,
  })), 1000);
}
```