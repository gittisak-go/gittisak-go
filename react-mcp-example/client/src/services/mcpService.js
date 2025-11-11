/**
 * MCP Service
 * 
 * This service handles communication with the MCP server.
 * In a production environment, this would use the MCP SDK to communicate
 * with the server over stdio or HTTP transport.
 * 
 * For this demo, we're simulating the server responses directly in the browser.
 */

// Simulated in-memory storage
let tasks = new Map()
let notes = new Map()

/**
 * Simulate MCP server tool call
 */
async function callTool(toolName, args) {
  // Simulate network delay
  await new Promise(resolve => setTimeout(resolve, 300))

  switch (toolName) {
    case 'create_task': {
      const taskId = `task_${Date.now()}`
      const task = {
        id: taskId,
        title: args.title,
        description: args.description || '',
        priority: args.priority || 'medium',
        status: 'pending',
        createdAt: new Date().toISOString(),
      }
      tasks.set(taskId, task)
      
      return {
        success: true,
        message: 'งานถูกสร้างเรียบร้อยแล้ว / Task created successfully',
        task,
      }
    }

    case 'list_tasks': {
      const status = args.status || 'all'
      let filteredTasks = Array.from(tasks.values())
      
      if (status !== 'all') {
        filteredTasks = filteredTasks.filter(task => task.status === status)
      }

      return {
        success: true,
        count: filteredTasks.length,
        tasks: filteredTasks,
      }
    }

    case 'complete_task': {
      const task = tasks.get(args.taskId)
      
      if (!task) {
        throw new Error('ไม่พบงาน / Task not found')
      }

      task.status = 'completed'
      task.completedAt = new Date().toISOString()
      tasks.set(args.taskId, task)

      return {
        success: true,
        message: 'งานถูกทำเครื่องหมายว่าเสร็จสิ้น / Task marked as completed',
        task,
      }
    }

    case 'create_note': {
      const noteId = `note_${Date.now()}`
      const note = {
        id: noteId,
        title: args.title,
        content: args.content,
        tags: args.tags || [],
        createdAt: new Date().toISOString(),
      }
      notes.set(noteId, note)

      return {
        success: true,
        message: 'โน้ตถูกสร้างเรียบร้อยแล้ว / Note created successfully',
        note,
      }
    }

    case 'search_notes': {
      const query = args.query.toLowerCase()
      const results = Array.from(notes.values()).filter(note => 
        note.title.toLowerCase().includes(query) ||
        note.content.toLowerCase().includes(query) ||
        note.tags.some(tag => tag.toLowerCase().includes(query))
      )

      return {
        success: true,
        count: results.length,
        notes: results,
      }
    }

    case 'get_weather': {
      // Simulated weather data
      const weatherData = {
        city: args.city,
        temperature: Math.floor(Math.random() * 15) + 20, // 20-35°C
        condition: ['Sunny', 'Cloudy', 'Rainy', 'Partly Cloudy'][Math.floor(Math.random() * 4)],
        humidity: Math.floor(Math.random() * 40) + 40, // 40-80%
        timestamp: new Date().toISOString(),
      }

      return {
        success: true,
        weather: weatherData,
      }
    }

    default:
      throw new Error('ไม่พบเครื่องมือนี้ / Tool not found')
  }
}

/**
 * List available MCP tools
 */
async function listTools() {
  // Simulate network delay
  await new Promise(resolve => setTimeout(resolve, 200))

  return [
    {
      name: 'create_task',
      description: 'สร้างงานใหม่ในระบบ / Create a new task',
    },
    {
      name: 'list_tasks',
      description: 'แสดงรายการงานทั้งหมด / List all tasks',
    },
    {
      name: 'complete_task',
      description: 'ทำเครื่องหมายงานว่าเสร็จสิ้น / Mark a task as completed',
    },
    {
      name: 'create_note',
      description: 'สร้างโน้ตใหม่ / Create a new note',
    },
    {
      name: 'search_notes',
      description: 'ค้นหาโน้ต / Search notes',
    },
    {
      name: 'get_weather',
      description: 'รับข้อมูลสภาพอากาศ / Get weather information',
    },
  ]
}

export const mcpService = {
  callTool,
  listTools,
}
