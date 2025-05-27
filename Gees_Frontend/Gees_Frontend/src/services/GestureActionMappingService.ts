import type { GestureActionMapping } from '@/models/gestureActionMapping.ts'

/**
 * Fetches all gesture action mappings from the API.
 * @returns {Promise<GestureActionMapping[]>} A promise that resolves to an array of GestureActionMapping objects.
 */
export async function updateGestureActionMappings(mappings: GestureActionMapping[]): Promise<void> {
  try {
    const endpoint = import.meta.env.VITE_API_URL
    const response = await fetch(endpoint + '/update-gesture-action-mappings', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(mappings),
    })
    if (!response.ok) {
      throw new Error('Network response was not ok')
    }
  } catch (error) {
    console.error('Error saving gesture mappings:', error)
  }
}

/**
 * Fetches all gesture action mappings from the API.
 */
export async function fetchGestureActionMappings(): Promise<GestureActionMapping[]> {
  try {
    const endpoint = import.meta.env.VITE_API_URL
    const response = await fetch(endpoint + '/gesture-action-mappings')
    if (!response.ok) {
      throw new Error('Network response was not ok')
    }
    return await response.json()
  } catch (error) {
    console.error('Error fetching gesture action mappings:', error)
    return []
  }
}
