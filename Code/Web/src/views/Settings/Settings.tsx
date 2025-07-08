import React, { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import i18n from 'i18next'

import { Segmented, Upload, UploadFile, UploadProps, message } from 'antd'
import {
  CheckCircleOutlined,
  CloseCircleOutlined,
  EditOutlined,
  LogoutOutlined
} from '@ant-design/icons'

import { useAuth } from '@/context/authContext'
import useImageCache from '@/hooks/Image/useImageCache'
import { Button, Input } from '@/components/common'
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
    <section
      className={style.settingsContainer}
      aria-labelledby="user-settings-title"
    >
      <h2 id="user-settings-title" className="sr-only">
        {t('pages.settings.user_settings_title')}
      </h2>
      <div className={style.userItem}>
        <Upload
          fileList={fileList}
          onChange={handleChange}
          beforeUpload={beforeUpload}
          showUploadList={false}
          listType="picture-circle"
          maxCount={1}
          customRequest={customRequest}
          aria-label={t('pages.settings.upload_picture_aria')}
        >
          <img
            src={isLoading ? DefaultUser : picture || DefaultUser}
            alt={t('pages.settings.profile_picture_alt')}
            className={style.image}
          />
        </Upload>
        {isUploading && (
          <div role="status" aria-live="polite" className="sr-only">
            {t('pages.settings.uploading_picture')}
          </div>
        )}
      </div>
      <div className={style.userInformations}>
        <div className={style.titleContainer}>
          <h3 id="user-info-title">{t('pages.settings.user_infos')}</h3>
          <div
            className={style.editButtons}
            role="toolbar"
            aria-label={t('pages.settings.edit_toolbar_aria')}
          >
            {editData && (
              <Button
                type="link"
                style={{ width: 35, height: 35, padding: 10 }}
                onClick={cancelEdit}
                aria-label={t('pages.settings.cancel_edit_aria')}
              >
                <CloseCircleOutlined
                  style={{ fontSize: '20px', color: 'red' }}
                  aria-hidden="true"
                />
              </Button>
            )}
            <Button
              type="link"
              style={{ width: 35, height: 35, padding: 10 }}
              onClick={() =>
                editData ? saveNewData() : setEditData(!editData)
              }
              aria-label={
                editData
                  ? t('pages.settings.save_changes_aria')
                  : t('pages.settings.edit_profile_aria')
              }
            >
              {!editData ? (
                <EditOutlined style={{ fontSize: '20px' }} aria-hidden="true" />
              ) : (
                <CheckCircleOutlined
                  style={{ fontSize: '20px', color: 'green' }}
                  aria-hidden="true"
                />
              )}
            </Button>
          </div>
        </div>
        <div role="group" aria-labelledby="user-info-title">
          <SubtitledElement
            subtitleKey={t('components.input.first_name.label')}
            subTitleStyle={{ opacity: 0.6 }}
          >
            {!editData ? (
              user?.firstname
            ) : (
              <Input
                className="input"
                defaultValue={user?.firstname}
                onChange={e => setNewData({ ...newData, firstname: e })}
                aria-label={t('components.input.first_name.label')}
                id="firstname-input"
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
                className="input"
                defaultValue={user?.lastname}
                onChange={e => setNewData({ ...newData, lastname: e })}
                aria-label={t('components.input.last_name.label')}
                id="lastname-input"
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
    </section>
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
      <main
        className={style.pageContainer}
        aria-labelledby="settings-page-title"
      >
        <header className={style.pageHeader}>
          <PageTitle
            title={t('pages.settings.title')}
            size="title"
            id="settings-page-title"
          />
        </header>
        <div className={style.layoutContainer}>
          <UserSettings t={t} />

          <section
            className={style.settingsContainer}
            aria-labelledby="app-settings-title"
          >
            <h2 id="app-settings-title" className="sr-only">
              {t('pages.settings.app_settings_title')}
            </h2>
            <div className={style.settingsItem}>
              <label htmlFor="language-selector" className={style.settingLabel}>
                {t('pages.settings.language')}
              </label>
              <Segmented
                id="language-selector"
                options={[
                  { label: t('pages.settings.fr'), value: 'fr' },
                  { label: t('pages.settings.en'), value: 'en' }
                ]}
                value={i18n.language}
                onChange={value => switchLanguage(value as string)}
                aria-label={t('pages.settings.language_selector_aria')}
                aria-describedby="language-help"
              />
              <div id="language-help" className="sr-only">
                {t('pages.settings.language_help')}
              </div>
            </div>

            <div className={style.settingsItem}>
              <span className={style.settingLabel}>
                {t('pages.settings.logout')}
              </span>
              <Button
                type="primary"
                danger
                shape="circle"
                icon={<LogoutOutlined aria-hidden="true" />}
                onClick={() => logout()}
                aria-label={t('pages.settings.logout_aria')}
                aria-describedby="logout-help"
              />
              <div id="logout-help" className="sr-only">
                {t('pages.settings.logout_help')}
              </div>
            </div>
          </section>
        </div>
      </main>
    </>
  )
}

export default Settings
