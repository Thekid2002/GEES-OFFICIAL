import { Gesture } from '@/models/gesture.ts'

/**
 * Fetches all gestures from the API.
 * @returns {Promise<Gesture[]>} A promise that resolves to an array of Gesture objects.
 */
export async function fetchGestures(): Promise<Gesture[]> {
  try {
    const endpoint = import.meta.env.VITE_API_URL
    const response = await fetch(endpoint + '/gestures')
    if (!response.ok) {
      throw new Error('Network response was not ok')
    }
    return await response.json()
  } catch (error) {
    console.error('Error fetching gestures:', error)
    return []
  }
}

/**
 * Fetches a gesture by ID from the API.
 * @param id The ID of the gesture to fetch.
 * @returns {Promise<Gesture | null>} A promise that resolves to the Gesture object or null if not found.
 */
export async function fetchGestureById(id: number): Promise<Gesture | null> {
  try {
    const endpoint = import.meta.env.VITE_API_URL
    const response = await fetch(endpoint + '/gesture/' + id)
    if (!response.ok) {
      throw new Error('Network response was not ok')
    }
    return await response.json()
  } catch (error) {
    console.error('Error fetching gesture:', error)
    return null
  }
}

/**
 * Updates a gesture in the API.
 * @param id The ID of the gesture to update.
 * @param gesture The Gesture object with updated data.
 * @returns {Promise<void>} A promise that resolves when the gesture is updated.
 */
export async function updateGestureById(id: number, gesture: Gesture): Promise<void> {
  try {
    const endpoint = import.meta.env.VITE_API_URL
    const response = await fetch(endpoint + '/gesture/' + id, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(gesture),
    })
    if (!response.ok) {
      throw new Error('Network response was not ok')
    }
  } catch (error) {
    console.error('Error updating gesture:', error)
  }
}

/**
 * Creates a new gesture in the API.
 * @param gesture The Gesture object to create.
 * @returns {Promise<void>} A promise that resolves when the gesture is created.
 */
export async function createGesture(gesture: Gesture): Promise<void> {
  try {
    const endpoint = import.meta.env.VITE_API_URL
    const response = await fetch(endpoint + '/gesture', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(gesture),
    })
    if (!response.ok) {
      throw new Error('Network response was not ok')
    }
  } catch (error) {
    console.error('Error creating gesture:', error)
  }
}
