import { useState } from 'react'
import { useMCP } from '../hooks/useMCP'

function NoteManager() {
  const { callTool, loading, error } = useMCP()
  const [notes, setNotes] = useState([])
  const [searchQuery, setSearchQuery] = useState('')
  const [formData, setFormData] = useState({
    title: '',
    content: '',
    tags: ''
  })

  const handleSubmit = async (e) => {
    e.preventDefault()
    
    const tagsArray = formData.tags
      .split(',')
      .map(tag => tag.trim())
      .filter(tag => tag.length > 0)

    const result = await callTool('create_note', {
      title: formData.title,
      content: formData.content,
      tags: tagsArray
    })
    
    if (result?.note) {
      setNotes([result.note, ...notes])
      setFormData({ title: '', content: '', tags: '' })
    }
  }

  const handleSearch = async (e) => {
    e.preventDefault()
    if (!searchQuery.trim()) return

    const result = await callTool('search_notes', { query: searchQuery })
    if (result?.notes) {
      setNotes(result.notes)
    }
  }

  return (
    <div className="note-manager">
      <div className="card">
        <h2>สร้างโน้ตใหม่ / Create New Note</h2>
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            placeholder="หัวข้อ / Title"
            value={formData.title}
            onChange={(e) => setFormData({ ...formData, title: e.target.value })}
            required
          />
          <textarea
            placeholder="เนื้อหา / Content"
            value={formData.content}
            onChange={(e) => setFormData({ ...formData, content: e.target.value })}
            rows="5"
            required
          />
          <input
            type="text"
            placeholder="แท็ก (คั่นด้วยจุลภาค) / Tags (comma separated)"
            value={formData.tags}
            onChange={(e) => setFormData({ ...formData, tags: e.target.value })}
          />
          <button type="submit" disabled={loading}>
            {loading ? 'กำลังสร้าง... / Creating...' : '➕ สร้างโน้ต / Create Note'}
          </button>
        </form>
      </div>

      <div className="card">
        <h2>ค้นหาโน้ต / Search Notes</h2>
        <form onSubmit={handleSearch}>
          <input
            type="text"
            placeholder="ค้นหา... / Search..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
          />
          <button type="submit" disabled={loading}>
            {loading ? 'กำลังค้นหา... / Searching...' : '🔍 ค้นหา / Search'}
          </button>
        </form>

        {error && <div className="error">{error}</div>}

        {notes.length === 0 ? (
          <p style={{ opacity: 0.6, marginTop: '1rem' }}>
            ยังไม่มีโน้ต / No notes yet
          </p>
        ) : (
          <div style={{ marginTop: '1rem' }}>
            {notes.map((note) => (
              <div key={note.id} className="note-item">
                <h3>{note.title}</h3>
                <p>{note.content}</p>
                {note.tags.length > 0 && (
                  <div style={{ marginTop: '0.5rem' }}>
                    {note.tags.map((tag, index) => (
                      <span key={index} className="tag">
                        {tag}
                      </span>
                    ))}
                  </div>
                )}
                <div style={{ marginTop: '0.5rem', fontSize: '0.875rem', opacity: 0.6 }}>
                  สร้างเมื่อ / Created: {new Date(note.createdAt).toLocaleString('th-TH')}
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  )
}

export default NoteManager
