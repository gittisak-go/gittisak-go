#!/usr/bin/env node

/**
 * React MCP Server Example
 * 
 * This is a production-ready MCP (Model Context Protocol) server implementation
 * that demonstrates how to create tools and resources for AI models.
 * 
 * Features:
 * - Task management tools
 * - Note-taking capabilities
 * - Weather information
 * - Resource management
 */

import { Server } from '@modelcontextprotocol/sdk/server/index.js';
import { StdioServerTransport } from '@modelcontextprotocol/sdk/server/stdio.js';
import {
  CallToolRequestSchema,
  ListToolsRequestSchema,
  ListResourcesRequestSchema,
  ReadResourceRequestSchema,
} from '@modelcontextprotocol/sdk/types.js';

// In-memory storage (in production, use a real database)
const tasks = new Map();
const notes = new Map();

/**
 * Initialize MCP Server
 */
const server = new Server(
  {
    name: 'react-mcp-example-server',
    version: '1.0.0',
  },
  {
    capabilities: {
      tools: {},
      resources: {},
    },
  }
);

/**
 * List available tools
 */
server.setRequestHandler(ListToolsRequestSchema, async () => {
  return {
    tools: [
      {
        name: 'create_task',
        description: 'สร้างงานใหม่ในระบบ / Create a new task',
        inputSchema: {
          type: 'object',
          properties: {
            title: {
              type: 'string',
              description: 'ชื่องาน / Task title',
            },
            description: {
              type: 'string',
              description: 'รายละเอียดงาน / Task description',
            },
            priority: {
              type: 'string',
              enum: ['low', 'medium', 'high'],
              description: 'ระดับความสำคัญ / Priority level',
            },
          },
          required: ['title'],
        },
      },
      {
        name: 'list_tasks',
        description: 'แสดงรายการงานทั้งหมด / List all tasks',
        inputSchema: {
          type: 'object',
          properties: {
            status: {
              type: 'string',
              enum: ['pending', 'completed', 'all'],
              description: 'กรองตามสถานะ / Filter by status',
            },
          },
        },
      },
      {
        name: 'complete_task',
        description: 'ทำเครื่องหมายงานว่าเสร็จสิ้น / Mark a task as completed',
        inputSchema: {
          type: 'object',
          properties: {
            taskId: {
              type: 'string',
              description: 'รหัสงาน / Task ID',
            },
          },
          required: ['taskId'],
        },
      },
      {
        name: 'create_note',
        description: 'สร้างโน้ตใหม่ / Create a new note',
        inputSchema: {
          type: 'object',
          properties: {
            title: {
              type: 'string',
              description: 'หัวข้อโน้ต / Note title',
            },
            content: {
              type: 'string',
              description: 'เนื้อหาโน้ต / Note content',
            },
            tags: {
              type: 'array',
              items: { type: 'string' },
              description: 'แท็ก / Tags',
            },
          },
          required: ['title', 'content'],
        },
      },
      {
        name: 'search_notes',
        description: 'ค้นหาโน้ต / Search notes',
        inputSchema: {
          type: 'object',
          properties: {
            query: {
              type: 'string',
              description: 'คำค้นหา / Search query',
            },
          },
          required: ['query'],
        },
      },
      {
        name: 'get_weather',
        description: 'รับข้อมูลสภาพอากาศ (ตัวอย่าง) / Get weather information (example)',
        inputSchema: {
          type: 'object',
          properties: {
            city: {
              type: 'string',
              description: 'ชื่อเมือง / City name',
            },
          },
          required: ['city'],
        },
      },
    ],
  };
});

/**
 * Handle tool calls
 */
server.setRequestHandler(CallToolRequestSchema, async (request) => {
  const { name, arguments: args } = request.params;

  try {
    switch (name) {
      case 'create_task': {
        const taskId = `task_${Date.now()}`;
        const task = {
          id: taskId,
          title: args.title,
          description: args.description || '',
          priority: args.priority || 'medium',
          status: 'pending',
          createdAt: new Date().toISOString(),
        };
        tasks.set(taskId, task);
        
        return {
          content: [
            {
              type: 'text',
              text: JSON.stringify({
                success: true,
                message: 'งานถูกสร้างเรียบร้อยแล้ว / Task created successfully',
                task,
              }, null, 2),
            },
          ],
        };
      }

      case 'list_tasks': {
        const status = args.status || 'all';
        let filteredTasks = Array.from(tasks.values());
        
        if (status !== 'all') {
          filteredTasks = filteredTasks.filter(task => task.status === status);
        }

        return {
          content: [
            {
              type: 'text',
              text: JSON.stringify({
                success: true,
                count: filteredTasks.length,
                tasks: filteredTasks,
              }, null, 2),
            },
          ],
        };
      }

      case 'complete_task': {
        const task = tasks.get(args.taskId);
        
        if (!task) {
          return {
            content: [
              {
                type: 'text',
                text: JSON.stringify({
                  success: false,
                  message: 'ไม่พบงาน / Task not found',
                }, null, 2),
              },
            ],
            isError: true,
          };
        }

        task.status = 'completed';
        task.completedAt = new Date().toISOString();
        tasks.set(args.taskId, task);

        return {
          content: [
            {
              type: 'text',
              text: JSON.stringify({
                success: true,
                message: 'งานถูกทำเครื่องหมายว่าเสร็จสิ้น / Task marked as completed',
                task,
              }, null, 2),
            },
          ],
        };
      }

      case 'create_note': {
        const noteId = `note_${Date.now()}`;
        const note = {
          id: noteId,
          title: args.title,
          content: args.content,
          tags: args.tags || [],
          createdAt: new Date().toISOString(),
        };
        notes.set(noteId, note);

        return {
          content: [
            {
              type: 'text',
              text: JSON.stringify({
                success: true,
                message: 'โน้ตถูกสร้างเรียบร้อยแล้ว / Note created successfully',
                note,
              }, null, 2),
            },
          ],
        };
      }

      case 'search_notes': {
        const query = args.query.toLowerCase();
        const results = Array.from(notes.values()).filter(note => 
          note.title.toLowerCase().includes(query) ||
          note.content.toLowerCase().includes(query) ||
          note.tags.some(tag => tag.toLowerCase().includes(query))
        );

        return {
          content: [
            {
              type: 'text',
              text: JSON.stringify({
                success: true,
                count: results.length,
                notes: results,
              }, null, 2),
            },
          ],
        };
      }

      case 'get_weather': {
        // Simulated weather data (in production, call a real weather API)
        const weatherData = {
          city: args.city,
          temperature: Math.floor(Math.random() * 15) + 20, // 20-35°C
          condition: ['Sunny', 'Cloudy', 'Rainy', 'Partly Cloudy'][Math.floor(Math.random() * 4)],
          humidity: Math.floor(Math.random() * 40) + 40, // 40-80%
          timestamp: new Date().toISOString(),
        };

        return {
          content: [
            {
              type: 'text',
              text: JSON.stringify({
                success: true,
                weather: weatherData,
              }, null, 2),
            },
          ],
        };
      }

      default:
        return {
          content: [
            {
              type: 'text',
              text: JSON.stringify({
                success: false,
                message: 'ไม่พบเครื่องมือนี้ / Tool not found',
              }, null, 2),
            },
          ],
          isError: true,
        };
    }
  } catch (error) {
    return {
      content: [
        {
          type: 'text',
          text: JSON.stringify({
            success: false,
            message: `ข้อผิดพลาด / Error: ${error.message}`,
          }, null, 2),
        },
      ],
      isError: true,
    };
  }
});

/**
 * List available resources
 */
server.setRequestHandler(ListResourcesRequestSchema, async () => {
  const taskResources = Array.from(tasks.values()).map(task => ({
    uri: `task:///${task.id}`,
    mimeType: 'application/json',
    name: task.title,
    description: task.description || 'No description',
  }));

  const noteResources = Array.from(notes.values()).map(note => ({
    uri: `note:///${note.id}`,
    mimeType: 'application/json',
    name: note.title,
    description: `Note with ${note.tags.length} tags`,
  }));

  return {
    resources: [...taskResources, ...noteResources],
  };
});

/**
 * Read a specific resource
 */
server.setRequestHandler(ReadResourceRequestSchema, async (request) => {
  const { uri } = request.params;

  if (uri.startsWith('task:///')) {
    const taskId = uri.replace('task:///', '');
    const task = tasks.get(taskId);

    if (!task) {
      throw new Error('Task not found');
    }

    return {
      contents: [
        {
          uri,
          mimeType: 'application/json',
          text: JSON.stringify(task, null, 2),
        },
      ],
    };
  }

  if (uri.startsWith('note:///')) {
    const noteId = uri.replace('note:///', '');
    const note = notes.get(noteId);

    if (!note) {
      throw new Error('Note not found');
    }

    return {
      contents: [
        {
          uri,
          mimeType: 'application/json',
          text: JSON.stringify(note, null, 2),
        },
      ],
    };
  }

  throw new Error('Unknown resource type');
});

/**
 * Start the server
 */
async function main() {
  const transport = new StdioServerTransport();
  await server.connect(transport);
  console.error('React MCP Example Server running on stdio');
}

main().catch((error) => {
  console.error('Fatal error in main():', error);
  process.exit(1);
});
