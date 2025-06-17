import React, { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import i18n from 'i18next'

import {
  Segmented,
  Upload,
  UploadFile,
  UploadProps,
  Input,
  message
} from 'antd'
import {
  CheckCircleOutlined,
  CloseCircleOutlined,
  EditOutlined,
  LogoutOutlined
} from '@ant-design/icons'

import { useAuth } from '@/context/authContext'
import useImageCache from '@/hooks/Image/useImageCache'
import { Button } from '@/components/common'
import SubtitledElement from '@/components/ui/SubtitledElement/SubtitledElement'
import PageMeta from '@/components/ui/PageMeta/PageMeta'
import PageTitle from '@/components/ui/PageText/Title'
import UpdateUserInfos from '@/services/api/User/UpdateUserInfos'
import GetUserPicture from '@/services/api/User/GetUserPicture'
import PutUserPicture from '@/services/api/User/PutUserPicture'

import DefaultUser from '@/assets/images/DefaultProfile.png'
import style from './Settings.module.css'

interface UserSettingsProps {
  t: (key: string) => string
}

const UserSettings: React.FC<UserSettingsProps> = ({ t }) => {
  const { user, updateUser } = useAuth()
  const [fileList, setFileList] = useState<UploadFile[]>([])
  const [isUploading, setIsUploading] = useState(false)

  const {
    data: picture,
    isLoading,
    updateCache
  } = useImageCache(user?.id, GetUserPicture)

  const putUserPicture = async (pictureData: string) => {
    if (isUploading) return

    setIsUploading(true)
    try {
      const req = await PutUserPicture(pictureData)
      if (req) {
        message.success(t('components.messages.picture_updated'))
        await updateCache(pictureData)
      } else {
        message.error(t('components.messages.picture_not_updated'))
      }
    } catch (error) {
      console.error('Error updating user picture:', error)
    } finally {
      setIsUploading(false)
    }
  }

  const handleChange: UploadProps['onChange'] = ({ fileList: newFileList }) => {
    setFileList(newFileList)
  }

  const customRequest: UploadProps['customRequest'] = async ({ file }) => {
    try {
      const reader = new FileReader()
      reader.onload = async e => {
        if (e.target?.result) {
          const base64 = e.target.result as string
          await putUserPicture(base64)
        }
      }
      reader.readAsDataURL(file as File)
    } catch (error) {
      message.error(t('components.messages.error_uploading_picture'))
    }
  }

  const beforeUpload = (file: File) => {
    const isImage = file.type.startsWith('image/')
    if (!isImage) {
      message.error('Only image files are allowed')
    }
    return isImage
  }

  const [editData, setEditData] = useState(false)
  const [oldData, setOldData] = useState({
    firstname: user?.firstname,
    lastname: user?.lastname,
    email: user?.email
  })
  const [newData, setNewData] = useState({
    firstname: user?.firstname,
    lastname: user?.lastname,
    email: user?.email
  })

  const saveNewData = async () => {
    if (
      newData.firstname === oldData.firstname &&
      newData.lastname === oldData.lastname
    ) {
      setEditData(false)
      message.info(t('components.messages.no_modifications'))
      return
    }
    try {
      await UpdateUserInfos({
        firstname: newData.firstname as string,
        lastname: newData.lastname as string
      })
      setOldData(newData)
      updateUser(newData)
      setEditData(false)
      message.success(t('components.messages.modifications_saved'))
    } catch (error) {
      console.error('Error updating user data:', error)
    }
  }

  const cancelEdit = () => {
    setNewData(oldData)
    setEditData(false)
  }

  useEffect(() => {
    if (user) {
      setNewData({
        firstname: user.firstname,
        lastname: user.lastname,
        email: user?.email
      })
      setOldData({
        firstname: user.firstname,
        lastname: user.lastname,
        email: user?.email
      })
    }
  }, [user])

  return (
    <div className={style.settingsContainer}>
      <div className={style.userItem}>
        <Upload
          fileList={fileList}
          onChange={handleChange}
          beforeUpload={beforeUpload}
          showUploadList={false}
          listType="picture-circle"
          maxCount={1}
          customRequest={customRequest}
        >
          <img
            src={isLoading ? DefaultUser : picture || DefaultUser}
            alt="user"
            className={style.image}
          />
        </Upload>
      </div>
      <div className={style.userInformations}>
        <div className={style.titleContainer}>
          <b>{t('pages.settings.user_infos')}</b>
          <div className={style.editButtons}>
            {editData && (
              <Button
                type="link"
                style={{ width: 35, height: 35, padding: 10 }}
                onClick={cancelEdit}
                aria-label={t('component.button.close')}
              >
                <CloseCircleOutlined
                  style={{ fontSize: '20px', color: 'red' }}
                />
              </Button>
            )}
            <Button
              type="link"
              style={{ width: 35, height: 35, padding: 10 }}
              onClick={() =>
                editData ? saveNewData() : setEditData(!editData)
              }
              aria-label={t('component.button.edit')}
            >
              {!editData ? (
                <EditOutlined style={{ fontSize: '20px' }} />
              ) : (
                <CheckCircleOutlined
                  style={{ fontSize: '20px', color: 'green' }}
                />
              )}
            </Button>
          </div>
        </div>
        <SubtitledElement
          subtitleKey={t('components.input.first_name.label')}
          subTitleStyle={{ opacity: 0.6 }}
        >
          {!editData ? (
            user?.firstname
          ) : (
            <Input
              defaultValue={user?.firstname}
              onChange={e =>
                setNewData({ ...newData, firstname: e.target.value })
              }
              aria-label={t('component.input.first_name.label')}
            />
          )}
        </SubtitledElement>
        <SubtitledElement
          subtitleKey={t('components.input.last_name.label')}
          subTitleStyle={{ opacity: 0.6 }}
        >
          {!editData ? (
            user?.lastname
          ) : (
            <Input
              defaultValue={user?.lastname}
              onChange={e =>
                setNewData({ ...newData, lastname: e.target.value })
              }
              aria-label={t('component.input.last_name.label')}
            />
          )}
        </SubtitledElement>
        <SubtitledElement
          subtitleKey={t('components.input.email.label')}
          subTitleStyle={{ opacity: 0.6 }}
        >
          {user?.email}
        </SubtitledElement>
      </div>
    </div>
  )
}

const Settings: React.FC = () => {
  const { t } = useTranslation()
  const { logout } = useAuth()

  const switchLanguage = (language: string) => {
    let lang = ''
    switch (language) {
      case 'fr' as string:
        lang = 'fr'
        localStorage.setItem('lang', 'fr')
        break
      case 'en' as string:
        lang = 'en'
        localStorage.setItem('lang', 'en')
        break
      default:
        lang = 'fr'
        localStorage.setItem('lang', 'fr')
        break
    }
    i18n.changeLanguage(lang)
  }

  return (
    <>
      <PageMeta
        title={t('pages.settings.document_title')}
        description={t('pages.settings.document_description')}
        keywords="settings, user, Keyz"
      />
      <div className={style.pageContainer}>
        <div className={style.pageHeader}>
          <PageTitle title={t('pages.settings.title')} size="title" />
        </div>
        <div className={style.layoutContainer}>
          <UserSettings t={t} />

          <div id="main-content" className={style.settingsContainer}>
            <div className={style.settingsItem}>
              {t('pages.settings.language')}
              <Segmented
                options={[
                  { label: t('pages.settings.fr'), value: 'fr' },
                  { label: t('pages.settings.en'), value: 'en' }
                ]}
                value={i18n.language}
                onChange={value => switchLanguage(value as string)}
                tabIndex={0}
              />
            </div>

            <div className={style.settingsItem}>
              {t('pages.settings.logout')}
              <Button
                type="primary"
                danger
                shape="circle"
                icon={<LogoutOutlined />}
                onClick={() => logout()}
              />
            </div>
          </div>
        </div>
      </div>
    </>
  )
}

export default Settings
