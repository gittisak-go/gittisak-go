import { useState } from 'react'
import { mcpService } from '../services/mcpService'

/**
 * Custom hook for MCP tool calls
 * This hook provides a simple interface for calling MCP tools from React components
 */
export function useMCP() {
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)

  const callTool = async (toolName, args = {}) => {
    setLoading(true)
    setError(null)

    try {
      const result = await mcpService.callTool(toolName, args)
      setLoading(false)
      return result
    } catch (err) {
      setError(err.message || 'เกิดข้อผิดพลาด / An error occurred')
      setLoading(false)
      return null
    }
  }

  const listTools = async () => {
    setLoading(true)
    setError(null)

    try {
      const tools = await mcpService.listTools()
      setLoading(false)
      return tools
    } catch (err) {
      setError(err.message || 'เกิดข้อผิดพลาด / An error occurred')
      setLoading(false)
      return []
    }
  }

  return {
    callTool,
    listTools,
    loading,
    error,
  }
}
