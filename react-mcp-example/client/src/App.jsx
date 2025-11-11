import { useState } from 'react'
import TaskManager from './components/TaskManager'
import NoteManager from './components/NoteManager'
import WeatherWidget from './components/WeatherWidget'
import './App.css'

function App() {
  const [activeTab, setActiveTab] = useState('tasks')

  return (
    <div className="App">
      <h1>🚀 React MCP Example</h1>
      <p className="subtitle">
        ตัวอย่าง Model Context Protocol Server พร้อม React Client<br/>
        Production-Ready MCP Server with React Client
      </p>

      <div className="tabs">
        <button 
          className={`tab ${activeTab === 'tasks' ? 'active' : ''}`}
          onClick={() => setActiveTab('tasks')}
        >
          📋 งาน / Tasks
        </button>
        <button 
          className={`tab ${activeTab === 'notes' ? 'active' : ''}`}
          onClick={() => setActiveTab('notes')}
        >
          📝 โน้ต / Notes
        </button>
        <button 
          className={`tab ${activeTab === 'weather' ? 'active' : ''}`}
          onClick={() => setActiveTab('weather')}
        >
          🌤️ สภาพอากาศ / Weather
        </button>
      </div>

      <div className="tab-content">
        {activeTab === 'tasks' && <TaskManager />}
        {activeTab === 'notes' && <NoteManager />}
        {activeTab === 'weather' && <WeatherWidget />}
      </div>

      <div className="info-card">
        <h3>ℹ️ เกี่ยวกับโปรเจกต์นี้ / About This Project</h3>
        <p>
          โปรเจกต์นี้แสดงตัวอย่างการใช้งาน MCP (Model Context Protocol) กับ React 
          ซึ่งสามารถนำไปใช้งานได้จริงในโปรเจกต์ Production
        </p>
        <p>
          This project demonstrates a production-ready example of using MCP 
          (Model Context Protocol) with React. It includes:
        </p>
        <ul>
          <li>✅ Task management with priorities</li>
          <li>✅ Note-taking with tags and search</li>
          <li>✅ Simulated weather information</li>
          <li>✅ MCP server implementation</li>
          <li>✅ React client with modern UI</li>
        </ul>
      </div>
    </div>
  )
}

export default App
