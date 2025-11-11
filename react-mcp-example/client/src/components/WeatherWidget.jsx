import { useState } from 'react'
import { useMCP } from '../hooks/useMCP'

function WeatherWidget() {
  const { callTool, loading, error } = useMCP()
  const [city, setCity] = useState('')
  const [weather, setWeather] = useState(null)

  const handleSubmit = async (e) => {
    e.preventDefault()
    if (!city.trim()) return

    const result = await callTool('get_weather', { city })
    if (result?.weather) {
      setWeather(result.weather)
    }
  }

  const getWeatherEmoji = (condition) => {
    const emojis = {
      'Sunny': '☀️',
      'Cloudy': '☁️',
      'Rainy': '🌧️',
      'Partly Cloudy': '⛅'
    }
    return emojis[condition] || '🌤️'
  }

  return (
    <div className="weather-widget">
      <div className="card">
        <h2>ตรวจสอบสภาพอากาศ / Check Weather</h2>
        <p style={{ opacity: 0.7, marginBottom: '1rem' }}>
          (ข้อมูลจำลอง / Simulated data for demonstration)
        </p>
        
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            placeholder="ชื่อเมือง / City name (e.g., Bangkok, Tokyo, London)"
            value={city}
            onChange={(e) => setCity(e.target.value)}
            required
          />
          <button type="submit" disabled={loading}>
            {loading ? 'กำลังโหลด... / Loading...' : '🌤️ ตรวจสอบ / Check Weather'}
          </button>
        </form>

        {error && <div className="error">{error}</div>}

        {weather && (
          <div className="weather-result" style={{
            marginTop: '2rem',
            padding: '2rem',
            background: 'linear-gradient(135deg, rgba(59, 130, 246, 0.2), rgba(147, 51, 234, 0.2))',
            borderRadius: '12px',
            textAlign: 'center'
          }}>
            <div style={{ fontSize: '4rem', marginBottom: '1rem' }}>
              {getWeatherEmoji(weather.condition)}
            </div>
            <h3 style={{ fontSize: '2rem', marginBottom: '1rem' }}>
              {weather.city}
            </h3>
            <div style={{ fontSize: '3rem', fontWeight: 'bold', marginBottom: '1rem' }}>
              {weather.temperature}°C
            </div>
            <div style={{ fontSize: '1.5rem', opacity: 0.9, marginBottom: '1rem' }}>
              {weather.condition}
            </div>
            <div style={{ opacity: 0.7 }}>
              💧 ความชื้น / Humidity: {weather.humidity}%
            </div>
            <div style={{ fontSize: '0.875rem', opacity: 0.5, marginTop: '1rem' }}>
              อัพเดต / Updated: {new Date(weather.timestamp).toLocaleString('th-TH')}
            </div>
          </div>
        )}
      </div>

      <div className="card">
        <h3>💡 เกี่ยวกับ Weather Tool</h3>
        <p>
          นี่คือตัวอย่างการสร้าง tool ใน MCP ที่สามารถเรียกใช้งานได้จาก AI model
          ในการใช้งานจริง คุณสามารถเชื่อมต่อกับ Weather API จริง เช่น OpenWeatherMap
        </p>
        <p style={{ marginTop: '1rem' }}>
          This demonstrates how to create an MCP tool that can be called by AI models.
          In production, you would connect to a real Weather API like OpenWeatherMap.
        </p>
      </div>
    </div>
  )
}

export default WeatherWidget
