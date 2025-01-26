import callApi from '@/services/api/apiCaller'

const PutUserPicture = async (pictureData: string) => {
  try {
    return await callApi({
      method: 'PUT',
      endpoint: `profile/picture`,
      data: JSON.stringify({ data: pictureData }),
    })
  } catch (error) {
    console.error('Error updating property picture:', error)
    throw error
  }
}

export default PutUserPicture
