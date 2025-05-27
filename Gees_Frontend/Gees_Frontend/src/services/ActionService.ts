import { Action } from '@/models/action.ts'

/**
 * Fetches an action by its ID from the API.
 * @param id The ID of the action to fetch.
 * @returns {Promise<Action | null>} A promise that resolves to the Action object if found, or null if not found.
 */
export async function fetchActionById(id: number): Promise<Action | null> {
  try {
    const endpoint = import.meta.env.VITE_API_URL
    const response = await fetch(endpoint + '/action/' + id)
    if (!response.ok) {
      throw new Error('Network response was not ok')
    }
    const action = await response.json()
    return action
  } catch (error) {
    console.error('Error fetching action:', error)
    return null
  }
}

/**
 * Updates an action by its ID in the API.
 * @param id The ID of the action to update.
 * @param action The Action object with updated data.
 */
export async function updateActionById(id: number, action: Action): Promise<void> {
  try {
    const endpoint = import.meta.env.VITE_API_URL
    const response = await fetch(endpoint + '/action/' + id, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(action),
    })
    if (!response.ok) {
      throw new Error('Network response was not ok')
    }
  } catch (error) {
    console.error('Error updating action:', error)
  }
}

/**
 * Creates a new action in the API.
 * @param action The Action object to create.
 * @returns {Promise<void>} A promise that resolves when the action is created.
 */
export async function createAction(action: Action): Promise<void> {
  try {
    const endpoint = import.meta.env.VITE_API_URL
    const response = await fetch(endpoint + '/action', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(action),
    })
    if (!response.ok) {
      throw new Error('Network response was not ok')
    }
  } catch (error) {
    console.error('Error creating action:', error)
  }
}

/**
 * Fetches all actions from the API.
 * @returns {Promise<Action[]>} A promise that resolves to an array of Action objects.
 */
export async function fetchActions(): Promise<Action[]> {
  try {
    const endpoint = import.meta.env.VITE_API_URL
    const response = await fetch(endpoint + '/actions')
    if (!response.ok) {
      throw new Error('Network response was not ok')
    }
    return await response.json()
  } catch (error) {
    console.error('Error fetching gestures:', error)
    return []
  }
}
