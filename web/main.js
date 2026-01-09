// K9s WASM initialization script

const go = new Go();

// Override console methods to capture output
const output = document.getElementById('output');
const statusDiv = document.getElementById('status');

const originalLog = console.log;
const originalError = console.error;
const originalWarn = console.warn;

console.log = function(...args) {
    originalLog.apply(console, args);
    appendOutput(args.join(' '));
};

console.error = function(...args) {
    originalError.apply(console, args);
    appendOutput('[ERROR] ' + args.join(' '), 'error');
};

console.warn = function(...args) {
    originalWarn.apply(console, args);
    appendOutput('[WARN] ' + args.join(' '), 'warn');
};

function appendOutput(text, type = 'info') {
    if (output) {
        output.style.display = 'block';
        const line = document.createElement('div');
        line.textContent = text;
        if (type === 'error') {
            line.style.color = '#ff4444';
        } else if (type === 'warn') {
            line.style.color = '#ffaa00';
        }
        output.appendChild(line);
        output.scrollTop = output.scrollHeight;
    }
}

function updateStatus(message, isError = false) {
    if (statusDiv) {
        const spinner = statusDiv.querySelector('.spinner');
        if (spinner) {
            spinner.style.display = isError ? 'none' : 'block';
        }
        
        const p = statusDiv.querySelector('p');
        if (p) {
            p.textContent = message;
            p.className = isError ? 'error' : 'success';
        }
    }
}

// Load and run the WASM module
WebAssembly.instantiateStreaming(fetch('k9s.wasm'), go.importObject)
    .then((result) => {
        updateStatus('Starting K9s...');
        go.run(result.instance);
        updateStatus('K9s is running in your browser!');
        appendOutput('=== K9s WebAssembly Demo Mode ===');
        appendOutput('Connected to demo-cluster with mock data');
        appendOutput('Use standard k9s keyboard shortcuts to navigate');
        appendOutput('Note: Some features are not available in demo mode');
        appendOutput('');
    })
    .catch((err) => {
        updateStatus('Failed to load K9s: ' + err.message, true);
        console.error('Failed to instantiate WASM module:', err);
        
        // Provide specific error messages for common issues
        let errorDetails = 'Error loading K9s: ' + err.message;
        if (err.message.includes('fetch')) {
            errorDetails += '\n\nNetwork Error: Unable to fetch k9s.wasm. Make sure:';
            errorDetails += '\n  • The web server is running';
            errorDetails += '\n  • k9s.wasm exists in the web directory';
            errorDetails += '\n  • CORS settings allow WASM file loading';
        } else if (err.message.includes('streaming')) {
            errorDetails += '\n\nStreaming Error: Your browser may not support WebAssembly streaming.';
            errorDetails += '\n  • Try a modern browser (Chrome, Firefox, Safari, Edge)';
            errorDetails += '\n  • Check that WASM is enabled in browser settings';
        } else if (err.message.includes('import')) {
            errorDetails += '\n\nImport Error: WASM import object mismatch.';
            errorDetails += '\n  • Ensure wasm_exec.js matches your Go version';
            errorDetails += '\n  • Try rebuilding with: make build-wasm';
        }
        
        appendOutput(errorDetails, 'error');
        appendOutput('', 'error');
        appendOutput('Make sure k9s.wasm and wasm_exec.js are present in the same directory.', 'error');
    });

// Handle WASM errors
window.addEventListener('error', (event) => {
    console.error('Global error:', event.error);
    appendOutput('Runtime error: ' + (event.error?.message || event.message), 'error');
});

window.addEventListener('unhandledrejection', (event) => {
    console.error('Unhandled promise rejection:', event.reason);
    appendOutput('Unhandled error: ' + (event.reason?.message || event.reason), 'error');
});
