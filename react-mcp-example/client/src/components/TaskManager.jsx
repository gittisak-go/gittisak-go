import { useState } from 'react'
import { useMCP } from '../hooks/useMCP'

function TaskManager() {
  const { callTool, loading, error } = useMCP()
  const [tasks, setTasks] = useState([])
  const [formData, setFormData] = useState({
    title: '',
    description: '',
    priority: 'medium'
  })

  const handleSubmit = async (e) => {
    e.preventDefault()
    
    const result = await callTool('create_task', formData)
    if (result?.task) {
      setTasks([result.task, ...tasks])
      setFormData({ title: '', description: '', priority: 'medium' })
    }
  }

  const loadTasks = async (status = 'all') => {
    const result = await callTool('list_tasks', { status })
    if (result?.tasks) {
      setTasks(result.tasks)
    }
  }

  const completeTask = async (taskId) => {
    const result = await callTool('complete_task', { taskId })
    if (result?.task) {
      setTasks(tasks.map(t => t.id === taskId ? result.task : t))
    }
  }

  return (
    <div className="task-manager">
      <div className="card">
        <h2>สร้างงานใหม่ / Create New Task</h2>
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            placeholder="ชื่องาน / Task title"
            value={formData.title}
            onChange={(e) => setFormData({ ...formData, title: e.target.value })}
            required
          />
          <textarea
            placeholder="รายละเอียด / Description (optional)"
            value={formData.description}
            onChange={(e) => setFormData({ ...formData, description: e.target.value })}
            rows="3"
          />
          <select
            value={formData.priority}
            onChange={(e) => setFormData({ ...formData, priority: e.target.value })}
          >
            <option value="low">ความสำคัญต่ำ / Low Priority</option>
            <option value="medium">ความสำคัญปานกลาง / Medium Priority</option>
            <option value="high">ความสำคัญสูง / High Priority</option>
          </select>
          <button type="submit" disabled={loading}>
            {loading ? 'กำลังสร้าง... / Creating...' : '➕ สร้างงาน / Create Task'}
          </button>
        </form>
      </div>

      <div className="card">
        <h2>รายการงาน / Task List</h2>
        <div style={{ marginBottom: '1rem' }}>
          <button onClick={() => loadTasks('all')} disabled={loading}>
            ทั้งหมด / All
          </button>
          <button onClick={() => loadTasks('pending')} disabled={loading} style={{ marginLeft: '0.5rem' }}>
            รอดำเนินการ / Pending
          </button>
          <button onClick={() => loadTasks('completed')} disabled={loading} style={{ marginLeft: '0.5rem' }}>
            เสร็จสิ้น / Completed
          </button>
        </div>

        {error && <div className="error">{error}</div>}

        {tasks.length === 0 ? (
          <p style={{ opacity: 0.6 }}>ยังไม่มีงาน / No tasks yet</p>
        ) : (
          tasks.map((task) => (
            <div 
              key={task.id} 
              className={`task-item ${task.status} ${task.priority}-priority`}
            >
              <h3>{task.title}</h3>
              {task.description && <p>{task.description}</p>}
              <div style={{ marginTop: '0.5rem', fontSize: '0.875rem', opacity: 0.8 }}>
                <span>ความสำคัญ / Priority: {task.priority.toUpperCase()}</span>
                {' • '}
                <span>สถานะ / Status: {task.status.toUpperCase()}</span>
              </div>
              {task.status === 'pending' && (
                <button 
                  onClick={() => completeTask(task.id)}
                  disabled={loading}
                  style={{ marginTop: '0.5rem' }}
                >
                  ✅ ทำเครื่องหมายเสร็จสิ้น / Mark Complete
                </button>
              )}
            </div>
          ))
        )}
      </div>
    </div>
  )
}

export default TaskManager
