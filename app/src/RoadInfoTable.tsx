import {
  Paper,
  Stack,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
} from '@mui/material'
import { FC } from 'react'

export const COLUMN_TEMPL_ROADWORK: Column[] = [
  {
    id: 'road',
    label: 'Road',
    align: 'left',
    externalKey: 'road',
  },
  {
    id: 'reason',
    label: 'Reason',
    align: 'left',
    externalKey: 'reason',
  },
]

export const COLUMN_TEMPL_JAMS: Column[] = [
  {
    id: 'road',
    label: 'Road',
    align: 'left',
    externalKey: 'road',
  },
  {
    id: 'distance',
    label: 'Distance',
    externalKey: 'distance',
  },
  {
    id: 'delay',
    label: 'Delay',
    externalKey: 'delay',
  },
  {
    id: 'reason',
    label: 'Reason',
    align: 'left',
    externalKey: 'reason',
  },
]

interface Column {
  id: string
  align?: 'center' | 'left' | 'right' | 'justify' | 'inherit'
  label: string
  externalKey: string
}

const toRow = (columnTempl: Column[], event: RoadEvent) => {
  return (
    <TableRow hover tabIndex={-1} key={event.id}>
      {columnTempl.map((column) => {
        const value = event[column.externalKey]
        return (
          <TableCell
            key={column.id}
            id={column.id}
            align={column.align}
            className="border px-2 py-2"
          >
            {value ?? '-'}
          </TableCell>
        )
      })}
    </TableRow>
  )
}

const RoadEventsTable: FC<{
  label: string
  columnTempl: Column[]
  data: RoadEvent[]
}> = ({ label, columnTempl, data: events }) => {
  return (
    <Stack direction={'column'}>
      <Typography>{label}</Typography>
      <TableContainer
        square
        component={Paper}
        sx={{ flex: 1, overflow: 'auto' }}
      >
        <Table stickyHeader aria-label="sticky table">
          <TableHead>
            <TableRow>
              {columnTempl.map((column) => (
                <TableCell align={column.align ?? 'center'} key={column.label}>
                  {column.label}
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>{events.map((e) => toRow(columnTempl, e))}</TableBody>
        </Table>
      </TableContainer>
    </Stack>
  )
}

export default RoadEventsTable
