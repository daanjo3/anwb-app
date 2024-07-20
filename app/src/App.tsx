import { useEffect, useMemo, useState } from 'react'
import './App.css'
import { Slider, Stack, Typography } from '@mui/material'
import { debounce } from '@mui/material/utils'
import RoadInfoTable, {
  COLUMN_TEMPL_JAMS,
  COLUMN_TEMPL_ROADWORK,
} from './RoadInfoTable'
import { Mark } from '@mui/material/Slider/useSlider.types'

const asMarks = (indexList: DocumentIndex[]): Mark[] =>
  indexList.map((index) => {
    const date = new Date(index._uploaded_at)
    return {
      value: date.getTime(),
    }
  })

function App() {
  const [jams, setJams] = useState<RoadEvent[]>([])
  const [roadworks, setRoadWorks] = useState<RoadEvent[]>([])
  const [indexList, setIndexList] = useState<DocumentIndex[]>([])
  const [header, setHeader] = useState('')
  const [selected, setSelected] = useState<DocumentIndex | undefined>(undefined)

  const getIndexByDate = (millis: number) =>
    indexList.find((index) => index._uploaded_at.getTime() == millis)

  // Fetch index on app render
  useEffect(() => {
    console.debug('Requesting index')
    fetch('http://localhost:8080/documents')
      .then((r) => r.json())
      .then((indexes) =>
        setIndexList(
          indexes.map((index: { id: string; _uploaded_at: string }) => ({
            id: index.id,
            _uploaded_at: new Date(index._uploaded_at),
          })),
        ),
      )
  }, [])

  // Set selected to latest on index load
  useEffect(() => {
    if (!indexList || indexList.length == 0) {
      return
    }
    setSelected(indexList[0]) // assume it's sorted
  }, [indexList])

  // Fetch road events on selection update
  useEffect(() => {
    if (!selected) {
      return
    }
    const newHeader = selected?._uploaded_at.toLocaleTimeString()
    if (newHeader != header) {
      // not sure if needed, want to avoid re-render
      setHeader(newHeader)
    }
    console.debug('Requesting jams')
    fetch(`http://localhost:8080/documents/${selected.id}/events/jams`)
      .then((r) => r.json())
      .then((jams) => setJams(jams))
    console.debug('Requesting roadworks')
    fetch(`http://localhost:8080/documents/${selected.id}/events/roadworks`)
      .then((r) => r.json())
      .then((roadworks) => setRoadWorks(roadworks))
  }, [selected])

  const marks = useMemo(() => asMarks(indexList), [indexList])
  const sliderOnChange = debounce((v: number) => {
    // unsure how this could be a number[] (but didn't check)
    const index = getIndexByDate(v)
    if (!index) {
      alert("Couldn't find index")
    }
    setSelected(index)
  }, 400)

  return (
    <>
      <Stack direction={'column'}>
        <Typography>{header}</Typography>
        {indexList.length > 0 && (
          <Slider
            aria-label="Time instances"
            defaultValue={marks[0].value}
            getAriaValueText={(v) =>
              getIndexByDate(v as number)?._uploaded_at.toLocaleDateString() ??
              '?'
            }
            max={marks[0].value}
            min={marks[marks.length - 1].value}
            step={null}
            valueLabelDisplay="auto"
            valueLabelFormat={(v) => new Date(v).toLocaleTimeString()}
            marks={marks}
            onChange={(_, v) => sliderOnChange(v as number)}
          />
        )}
        <Stack direction={'row'} spacing={1}>
          <RoadInfoTable
            label="Traffic jams"
            columnTempl={COLUMN_TEMPL_JAMS}
            data={jams}
          />
          <RoadInfoTable
            label="Road works"
            columnTempl={COLUMN_TEMPL_ROADWORK}
            data={roadworks}
          />
        </Stack>
      </Stack>
    </>
  )
}

export default App
