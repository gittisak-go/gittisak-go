# Code Examples / ตัวอย่างโค้ด

## Table of Contents
- [Basic Usage](#basic-usage)
- [Advanced Features](#advanced-features)
- [Custom Hooks](#custom-hooks)
- [Integration Examples](#integration-examples)
- [Error Handling](#error-handling)

---

## Basic Usage

### Example 1: Creating a Task

```javascript
import { useMCP } from '../hooks/useMCP'

function TaskForm() {
  const { callTool, loading, error } = useMCP()

  const handleSubmit = async (e) => {
    e.preventDefault()
    
    const result = await callTool('create_task', {
      title: 'Complete project documentation',
      description: 'Write comprehensive README and API docs',
      priority: 'high'
    })
    
    if (result?.task) {
      console.log('Task created:', result.task)
      // Handle success
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      {/* Form fields */}
      <button type="submit" disabled={loading}>
        {loading ? 'Creating...' : 'Create Task'}
      </button>
      {error && <div className="error">{error}</div>}
    </form>
  )
}
```

### Example 2: Listing Tasks with Filters

```javascript
function TaskList() {
  const { callTool, loading } = useMCP()
  const [tasks, setTasks] = useState([])

  const loadTasks = async (status = 'all') => {
    const result = await callTool('list_tasks', { status })
    if (result?.tasks) {
      setTasks(result.tasks)
    }
  }

  useEffect(() => {
    loadTasks()
  }, [])

  return (
    <div>
      <button onClick={() => loadTasks('pending')}>Pending</button>
      <button onClick={() => loadTasks('completed')}>Completed</button>
      <button onClick={() => loadTasks('all')}>All</button>
      
      {tasks.map(task => (
        <TaskItem key={task.id} task={task} />
      ))}
    </div>
  )
}
```

### Example 3: Creating and Searching Notes

```javascript
function NoteApp() {
  const { callTool } = useMCP()
  const [notes, setNotes] = useState([])

  // Create a note
  const createNote = async (title, content, tags) => {
    const result = await callTool('create_note', {
      title,
      content,
      tags: tags.split(',').map(t => t.trim())
    })
    
    if (result?.note) {
      setNotes([result.note, ...notes])
    }
  }

  // Search notes
  const searchNotes = async (query) => {
    const result = await callTool('search_notes', { query })
    if (result?.notes) {
      setNotes(result.notes)
    }
  }

  return (
    <div>
      <NoteForm onSubmit={createNote} />
      <SearchBar onSearch={searchNotes} />
      <NoteList notes={notes} />
    </div>
  )
}
```

---

## Advanced Features

### Example 4: Custom Hook with Caching

```javascript
import { useState, useCallback, useRef } from 'react'
import { mcpService } from '../services/mcpService'

export function useMCPWithCache() {
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)
  const cache = useRef(new Map())

  const callTool = useCallback(async (toolName, args = {}) => {
    // Create cache key
    const cacheKey = `${toolName}-${JSON.stringify(args)}`
    
    // Check cache
    if (cache.current.has(cacheKey)) {
      const cached = cache.current.get(cacheKey)
      const age = Date.now() - cached.timestamp
      
      // Return cached if less than 5 minutes old
      if (age < 5 * 60 * 1000) {
        return cached.data
      }
    }

    setLoading(true)
    setError(null)

    try {
      const result = await mcpService.callTool(toolName, args)
      
      // Store in cache
      cache.current.set(cacheKey, {
        data: result,
        timestamp: Date.now()
      })
      
      setLoading(false)
      return result
    } catch (err) {
      setError(err.message)
      setLoading(false)
      return null
    }
  }, [])

  const clearCache = useCallback(() => {
    cache.current.clear()
  }, [])

  return { callTool, loading, error, clearCache }
}
```

### Example 5: Batch Operations

```javascript
function BatchTaskCreator() {
  const { callTool } = useMCP()

  const createMultipleTasks = async (taskList) => {
    const results = await Promise.all(
      taskList.map(task => 
        callTool('create_task', task)
      )
    )
    
    const successful = results.filter(r => r?.success)
    const failed = results.filter(r => !r?.success)
    
    return {
      successful: successful.length,
      failed: failed.length,
      results
    }
  }

  const handleBatchCreate = async () => {
    const tasks = [
      { title: 'Task 1', priority: 'high' },
      { title: 'Task 2', priority: 'medium' },
      { title: 'Task 3', priority: 'low' }
    ]
    
    const result = await createMultipleTasks(tasks)
    console.log(`Created ${result.successful} tasks, ${result.failed} failed`)
  }

  return (
    <button onClick={handleBatchCreate}>
      Create Multiple Tasks
    </button>
  )
}
```

### Example 6: Real-time Updates with Polling

```javascript
function RealtimeTaskList() {
  const { callTool } = useMCP()
  const [tasks, setTasks] = useState([])
  const [isPolling, setIsPolling] = useState(false)

  useEffect(() => {
    if (!isPolling) return

    const interval = setInterval(async () => {
      const result = await callTool('list_tasks', { status: 'all' })
      if (result?.tasks) {
        setTasks(result.tasks)
      }
    }, 5000) // Poll every 5 seconds

    return () => clearInterval(interval)
  }, [isPolling, callTool])

  return (
    <div>
      <button onClick={() => setIsPolling(!isPolling)}>
        {isPolling ? 'Stop Updates' : 'Start Real-time Updates'}
      </button>
      <TaskList tasks={tasks} />
    </div>
  )
}
```

---

## Custom Hooks

### Example 7: useTask Hook

```javascript
function useTask() {
  const { callTool, loading, error } = useMCP()
  const [tasks, setTasks] = useState([])

  const createTask = useCallback(async (taskData) => {
    const result = await callTool('create_task', taskData)
    if (result?.task) {
      setTasks(prev => [result.task, ...prev])
      return result.task
    }
    return null
  }, [callTool])

  const loadTasks = useCallback(async (status = 'all') => {
    const result = await callTool('list_tasks', { status })
    if (result?.tasks) {
      setTasks(result.tasks)
    }
  }, [callTool])

  const completeTask = useCallback(async (taskId) => {
    const result = await callTool('complete_task', { taskId })
    if (result?.task) {
      setTasks(prev => 
        prev.map(t => t.id === taskId ? result.task : t)
      )
      return result.task
    }
    return null
  }, [callTool])

  const deleteTask = useCallback((taskId) => {
    setTasks(prev => prev.filter(t => t.id !== taskId))
  }, [])

  return {
    tasks,
    createTask,
    loadTasks,
    completeTask,
    deleteTask,
    loading,
    error
  }
}

// Usage
function TaskManager() {
  const {
    tasks,
    createTask,
    loadTasks,
    completeTask,
    loading,
    error
  } = useTask()

  useEffect(() => {
    loadTasks()
  }, [loadTasks])

  return (
    <div>
      {/* Use the hook's functions */}
    </div>
  )
}
```

### Example 8: useNote Hook

```javascript
function useNote() {
  const { callTool, loading, error } = useMCP()
  const [notes, setNotes] = useState([])

  const createNote = useCallback(async (noteData) => {
    const result = await callTool('create_note', noteData)
    if (result?.note) {
      setNotes(prev => [result.note, ...prev])
      return result.note
    }
    return null
  }, [callTool])

  const searchNotes = useCallback(async (query) => {
    const result = await callTool('search_notes', { query })
    if (result?.notes) {
      setNotes(result.notes)
    }
  }, [callTool])

  const filterByTag = useCallback((tag) => {
    return notes.filter(note => 
      note.tags.includes(tag)
    )
  }, [notes])

  return {
    notes,
    createNote,
    searchNotes,
    filterByTag,
    loading,
    error
  }
}
```

---

## Integration Examples

### Example 9: Integration with Redux

```javascript
// actions/mcpActions.js
export const createTask = (taskData) => async (dispatch) => {
  dispatch({ type: 'CREATE_TASK_REQUEST' })
  
  try {
    const result = await mcpService.callTool('create_task', taskData)
    
    if (result?.task) {
      dispatch({
        type: 'CREATE_TASK_SUCCESS',
        payload: result.task
      })
    }
  } catch (error) {
    dispatch({
      type: 'CREATE_TASK_FAILURE',
      error: error.message
    })
  }
}

// reducers/taskReducer.js
const initialState = {
  tasks: [],
  loading: false,
  error: null
}

export default function taskReducer(state = initialState, action) {
  switch (action.type) {
    case 'CREATE_TASK_REQUEST':
      return { ...state, loading: true, error: null }
    
    case 'CREATE_TASK_SUCCESS':
      return {
        ...state,
        loading: false,
        tasks: [action.payload, ...state.tasks]
      }
    
    case 'CREATE_TASK_FAILURE':
      return { ...state, loading: false, error: action.error }
    
    default:
      return state
  }
}
```

### Example 10: Integration with React Query

```javascript
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { mcpService } from '../services/mcpService'

// Query for listing tasks
function useTasks(status = 'all') {
  return useQuery({
    queryKey: ['tasks', status],
    queryFn: () => mcpService.callTool('list_tasks', { status }),
    select: (data) => data?.tasks || []
  })
}

// Mutation for creating task
function useCreateTask() {
  const queryClient = useQueryClient()
  
  return useMutation({
    mutationFn: (taskData) => 
      mcpService.callTool('create_task', taskData),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] })
    }
  })
}

// Usage
function TaskComponent() {
  const { data: tasks, isLoading } = useTasks('all')
  const createTask = useCreateTask()

  const handleCreate = async (taskData) => {
    await createTask.mutateAsync(taskData)
  }

  if (isLoading) return <div>Loading...</div>

  return (
    <div>
      <button onClick={() => handleCreate({ title: 'New Task' })}>
        Create Task
      </button>
      {tasks.map(task => (
        <TaskItem key={task.id} task={task} />
      ))}
    </div>
  )
}
```

---

## Error Handling

### Example 11: Comprehensive Error Handling

```javascript
function TaskManagerWithErrorHandling() {
  const { callTool } = useMCP()
  const [error, setError] = useState(null)

  const handleCreateTask = async (taskData) => {
    setError(null)
    
    try {
      const result = await callTool('create_task', taskData)
      
      if (!result) {
        throw new Error('No response from server')
      }
      
      if (!result.success) {
        throw new Error(result.message || 'Failed to create task')
      }
      
      // Success handling
      console.log('Task created successfully:', result.task)
      
    } catch (err) {
      // Error handling
      const errorMessage = err.message || 'An unexpected error occurred'
      setError(errorMessage)
      
      // Log error for debugging
      console.error('Error creating task:', err)
      
      // Optional: Send to error tracking service
      // trackError(err)
    }
  }

  return (
    <div>
      {error && (
        <div className="error-banner">
          <span>❌ {error}</span>
          <button onClick={() => setError(null)}>Dismiss</button>
        </div>
      )}
      {/* Rest of component */}
    </div>
  )
}
```

### Example 12: Retry Logic

```javascript
async function callToolWithRetry(toolName, args, maxRetries = 3) {
  let lastError

  for (let i = 0; i < maxRetries; i++) {
    try {
      const result = await mcpService.callTool(toolName, args)
      return result
    } catch (error) {
      lastError = error
      
      // Wait before retrying (exponential backoff)
      const delay = Math.pow(2, i) * 1000
      await new Promise(resolve => setTimeout(resolve, delay))
    }
  }

  throw new Error(`Failed after ${maxRetries} retries: ${lastError.message}`)
}

// Usage
function RobustTaskCreator() {
  const handleCreate = async (taskData) => {
    try {
      const result = await callToolWithRetry('create_task', taskData)
      console.log('Task created:', result)
    } catch (error) {
      console.error('Failed to create task:', error)
    }
  }

  return <button onClick={() => handleCreate({ title: 'Test' })}>
    Create Task
  </button>
}
```

---

## Testing Examples

### Example 13: Unit Tests

```javascript
// __tests__/useMCP.test.js
import { renderHook, act } from '@testing-library/react'
import { useMCP } from '../hooks/useMCP'
import { mcpService } from '../services/mcpService'

jest.mock('../services/mcpService')

describe('useMCP', () => {
  it('should call tool successfully', async () => {
    const mockResult = { success: true, task: { id: '1', title: 'Test' } }
    mcpService.callTool.mockResolvedValue(mockResult)

    const { result } = renderHook(() => useMCP())

    let response
    await act(async () => {
      response = await result.current.callTool('create_task', { title: 'Test' })
    })

    expect(response).toEqual(mockResult)
    expect(result.current.loading).toBe(false)
    expect(result.current.error).toBe(null)
  })

  it('should handle errors', async () => {
    mcpService.callTool.mockRejectedValue(new Error('Test error'))

    const { result } = renderHook(() => useMCP())

    await act(async () => {
      await result.current.callTool('create_task', { title: 'Test' })
    })

    expect(result.current.error).toBe('Test error')
    expect(result.current.loading).toBe(false)
  })
})
```

---

For more examples, check the source code in the `client/src/` directory.
