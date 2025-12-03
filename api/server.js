const express = require('express');
const app = express();
app.use(express.json());

// Mock endpoints for testing
app.get('/health', (req, res) => {
    res.json({ status: 'OK', message: 'Blue Carbon MRV API Running' });
});

app.post('/api/projects', (req, res) => {
    console.log('Project creation request:', req.body);
    res.json({
        success: true,
        message: 'Project created (mock)',
        projectId: 'PROJ_' + Date.now()
    });
});

app.get('/api/projects/:id', (req, res) => {
    res.json({
        id: req.params.id,
        name: 'Sample Mangrove Project',
        area: '50 hectares',const express = require('express');
const app = express();
app.use(express.json());

// Mock endpoints for testing
app.get('/health', (req, res) => {
    res.json({ status: 'OK', message: 'Blue Carbon MRV API Running' });
});

app.post('/api/projects', (req, res) => {
    console.log('Project creation request:', req.body);
    res.json({
        success: true,
        message: 'Project created (mock)',
        projectId: 'PROJ_' + Date.now()
    });
});

app.get('/api/projects/:id', (req, res) => {
    res.json({
        id: req.params.id,
        name: 'Sample Mangrove Project',
        area: '50 hectares',
        location: 'Sundarbans, India'
    });
});

const PORT = 3000;
app.listen(PORT, () => {
    console.log(`Server running on http://localhost:${PORT}`);
    console.log(`Health check: http://localhost:${PORT}/health`);
});
        location: 'Sundarbans, India'
    });
});

const PORT = 3000;
app.listen(PORT, () => {
    console.log(`Server running on http://localhost:${PORT}`);
    console.log(`Health check: http://localhost:${PORT}/health`);
});
