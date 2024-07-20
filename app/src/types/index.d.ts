/* eslint-disable @typescript-eslint/no-explicit-any */
interface RoadEvent extends Record<string, any> {
  id: number
  road: string
  distance?: number
  delay?: number
}

interface DocumentIndex extends Record<string, any> {
  id: string
  _uploaded_at: Date
}
