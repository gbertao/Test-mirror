// Copyright (C) 2022 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package selinux

var EXCEPTION_NO_FUZZER = []string{}

//
// To add a fuzzer for service, add your service name and fuzzer name in ServiceFuzzerBindings
// example of entry -
//	"android.hardware.health.IHealth/default": []string{"android.hardware.health-service.aidl_fuzzer"},

var (
	ServiceFuzzerBindings = map[string][]string{
		"android.hardware.audio.core.IConfig/default":                             EXCEPTION_NO_FUZZER,
		"android.hardware.audio.core.IModule/default":                             EXCEPTION_NO_FUZZER,
		"android.hardware.audio.core.IModule/a2dp":                                EXCEPTION_NO_FUZZER,
		"android.hardware.audio.core.IModule/bluetooth":                           EXCEPTION_NO_FUZZER,
		"android.hardware.audio.core.IModule/hearing_aid":                         EXCEPTION_NO_FUZZER,
		"android.hardware.audio.core.IModule/msd":                                 EXCEPTION_NO_FUZZER,
		"android.hardware.audio.core.IModule/r_submix":                            EXCEPTION_NO_FUZZER,
		"android.hardware.audio.core.IModule/stub":                                EXCEPTION_NO_FUZZER,
		"android.hardware.audio.core.IModule/usb":                                 EXCEPTION_NO_FUZZER,
		"android.hardware.audio.effect.IFactory/default":                          EXCEPTION_NO_FUZZER,
		"android.hardware.audio.sounddose.ISoundDoseFactory/default":              EXCEPTION_NO_FUZZER,
		"android.hardware.authsecret.IAuthSecret/default":                         EXCEPTION_NO_FUZZER,
		"android.hardware.automotive.evs.IEvsEnumerator/hw/0":                     EXCEPTION_NO_FUZZER,
		"android.hardware.boot.IBootControl/default":                              EXCEPTION_NO_FUZZER,
		"android.hardware.automotive.can.ICanController/default":                  EXCEPTION_NO_FUZZER,
		"android.hardware.automotive.evs.IEvsEnumerator/hw/1":                     EXCEPTION_NO_FUZZER,
		"android.hardware.automotive.ivn.IIvnAndroidDevice/default":               EXCEPTION_NO_FUZZER,
		"android.hardware.automotive.remoteaccess.IRemoteAccess/default":          EXCEPTION_NO_FUZZER,
		"android.hardware.automotive.vehicle.IVehicle/default":                    EXCEPTION_NO_FUZZER,
		"android.hardware.automotive.audiocontrol.IAudioControl/default":          EXCEPTION_NO_FUZZER,
		"android.hardware.biometrics.face.IFace/default":                          EXCEPTION_NO_FUZZER,
		"android.hardware.biometrics.face.IFace/virtual":                          EXCEPTION_NO_FUZZER,
		"android.hardware.biometrics.fingerprint.IFingerprint/default":            EXCEPTION_NO_FUZZER,
		"android.hardware.biometrics.fingerprint.IFingerprint/virtual":            EXCEPTION_NO_FUZZER,
		"android.hardware.bluetooth.audio.IBluetoothAudioProviderFactory/default": EXCEPTION_NO_FUZZER,
		"android.hardware.broadcastradio.IBroadcastRadio/amfm":                    EXCEPTION_NO_FUZZER,
		"android.hardware.broadcastradio.IBroadcastRadio/dab":                     EXCEPTION_NO_FUZZER,
		"android.hardware.bluetooth.IBluetoothHci/default":                        EXCEPTION_NO_FUZZER,
		"android.hardware.bluetooth.finder.IBluetoothFinder/default":              EXCEPTION_NO_FUZZER,
		"android.hardware.bluetooth.lmp_event.IBluetoothLmpEvent/default":         EXCEPTION_NO_FUZZER,
		"android.hardware.camera.provider.ICameraProvider/internal/0":             EXCEPTION_NO_FUZZER,
		"android.hardware.camera.provider.ICameraProvider/virtual/0":              EXCEPTION_NO_FUZZER,
		"android.hardware.cas.IMediaCasService/default":                           EXCEPTION_NO_FUZZER,
		"android.hardware.confirmationui.IConfirmationUI/default":                 []string{"android.hardware.confirmationui-service.trusty_fuzzer"},
		"android.hardware.contexthub.IContextHub/default":                         EXCEPTION_NO_FUZZER,
		"android.hardware.drm.IDrmFactory/clearkey":                               EXCEPTION_NO_FUZZER,
		"android.hardware.drm.ICryptoFactory/clearkey":                            EXCEPTION_NO_FUZZER,
		"android.hardware.dumpstate.IDumpstateDevice/default":                     EXCEPTION_NO_FUZZER,
		"android.hardware.fastboot.IFastboot/default":                             EXCEPTION_NO_FUZZER,
		"android.hardware.gatekeeper.IGatekeeper/default":                         EXCEPTION_NO_FUZZER,
		"android.hardware.gnss.IGnss/default":                                     EXCEPTION_NO_FUZZER,
		"android.hardware.graphics.allocator.IAllocator/default":                  EXCEPTION_NO_FUZZER,
		"android.hardware.graphics.composer3.IComposer/default":                   EXCEPTION_NO_FUZZER,
		"android.hardware.health.storage.IStorage/default":                        EXCEPTION_NO_FUZZER,
		"android.hardware.health.IHealth/default":                                 []string{"android.hardware.health-service.aidl_fuzzer"},
		"android.hardware.identity.IIdentityCredentialStore/default":              EXCEPTION_NO_FUZZER,
		"android.hardware.input.processor.IInputProcessor/default":                EXCEPTION_NO_FUZZER,
		"android.hardware.ir.IConsumerIr/default":                                 EXCEPTION_NO_FUZZER,
		"android.hardware.light.ILights/default":                                  EXCEPTION_NO_FUZZER,
		"android.hardware.macsec.IMacsecPskPlugin/default":                        EXCEPTION_NO_FUZZER,
		"android.hardware.media.c2.IComponentStore/default":                       EXCEPTION_NO_FUZZER,
		"android.hardware.media.c2.IComponentStore/software":                      []string{"libcodec2-aidl-fuzzer"},
		"android.hardware.memtrack.IMemtrack/default":                             EXCEPTION_NO_FUZZER,
		"android.hardware.net.nlinterceptor.IInterceptor/default":                 EXCEPTION_NO_FUZZER,
		"android.hardware.nfc.INfc/default":                                       EXCEPTION_NO_FUZZER,
		"android.hardware.oemlock.IOemLock/default":                               EXCEPTION_NO_FUZZER,
		"android.hardware.power.IPower/default":                                   EXCEPTION_NO_FUZZER,
		"android.hardware.power.stats.IPowerStats/default":                        EXCEPTION_NO_FUZZER,
		"android.hardware.radio.config.IRadioConfig/default":                      EXCEPTION_NO_FUZZER,
		"android.hardware.radio.data.IRadioData/slot1":                            EXCEPTION_NO_FUZZER,
		"android.hardware.radio.data.IRadioData/slot2":                            EXCEPTION_NO_FUZZER,
		"android.hardware.radio.data.IRadioData/slot3":                            EXCEPTION_NO_FUZZER,
		"android.hardware.radio.ims.IRadioIms/slot1":                              EXCEPTION_NO_FUZZER,
		"android.hardware.radio.ims.IRadioIms/slot2":                              EXCEPTION_NO_FUZZER,
		"android.hardware.radio.ims.IRadioIms/slot3":                              EXCEPTION_NO_FUZZER,
		"android.hardware.radio.ims.media.IImsMedia/default":                      EXCEPTION_NO_FUZZER,
		"android.hardware.radio.messaging.IRadioMessaging/slot1":                  EXCEPTION_NO_FUZZER,
		"android.hardware.radio.messaging.IRadioMessaging/slot2":                  EXCEPTION_NO_FUZZER,
		"android.hardware.radio.messaging.IRadioMessaging/slot3":                  EXCEPTION_NO_FUZZER,
		"android.hardware.radio.modem.IRadioModem/slot1":                          EXCEPTION_NO_FUZZER,
		"android.hardware.radio.modem.IRadioModem/slot2":                          EXCEPTION_NO_FUZZER,
		"android.hardware.radio.modem.IRadioModem/slot3":                          EXCEPTION_NO_FUZZER,
		"android.hardware.radio.network.IRadioNetwork/slot1":                      EXCEPTION_NO_FUZZER,
		"android.hardware.radio.network.IRadioNetwork/slot2":                      EXCEPTION_NO_FUZZER,
		"android.hardware.radio.network.IRadioNetwork/slot3":                      EXCEPTION_NO_FUZZER,
		"android.hardware.radio.satellite.IRadioSatellite/slot1":                  EXCEPTION_NO_FUZZER,
		"android.hardware.radio.satellite.IRadioSatellite/slot2":                  EXCEPTION_NO_FUZZER,
		"android.hardware.radio.satellite.IRadioSatellite/slot3":                  EXCEPTION_NO_FUZZER,
		"android.hardware.radio.sim.IRadioSim/slot1":                              EXCEPTION_NO_FUZZER,
		"android.hardware.radio.sim.IRadioSim/slot2":                              EXCEPTION_NO_FUZZER,
		"android.hardware.radio.sim.IRadioSim/slot3":                              EXCEPTION_NO_FUZZER,
		"android.hardware.radio.sap.ISap/slot1":                                   EXCEPTION_NO_FUZZER,
		"android.hardware.radio.sap.ISap/slot2":                                   EXCEPTION_NO_FUZZER,
		"android.hardware.radio.sap.ISap/slot3":                                   EXCEPTION_NO_FUZZER,
		"android.hardware.radio.voice.IRadioVoice/slot1":                          EXCEPTION_NO_FUZZER,
		"android.hardware.radio.voice.IRadioVoice/slot2":                          EXCEPTION_NO_FUZZER,
		"android.hardware.radio.voice.IRadioVoice/slot3":                          EXCEPTION_NO_FUZZER,
		"android.hardware.rebootescrow.IRebootEscrow/default":                     EXCEPTION_NO_FUZZER,
		"android.hardware.secure_element.ISecureElement/eSE1":                     EXCEPTION_NO_FUZZER,
		"android.hardware.secure_element.ISecureElement/eSE2":                     EXCEPTION_NO_FUZZER,
		"android.hardware.secure_element.ISecureElement/eSE3":                     EXCEPTION_NO_FUZZER,
		"android.hardware.secure_element.ISecureElement/SIM1":                     EXCEPTION_NO_FUZZER,
		"android.hardware.secure_element.ISecureElement/SIM2":                     EXCEPTION_NO_FUZZER,
		"android.hardware.secure_element.ISecureElement/SIM3":                     EXCEPTION_NO_FUZZER,
		"android.hardware.security.authgraph.IAuthGraphKeyExchange/nonsecure":     []string{"android.hardware.authgraph-service.nonsecure_fuzzer"},
		"android.hardware.security.dice.IDiceDevice/default":                      EXCEPTION_NO_FUZZER,
		"android.hardware.security.keymint.IKeyMintDevice/default":                EXCEPTION_NO_FUZZER,
		"android.hardware.security.keymint.IRemotelyProvisionedComponent/default": EXCEPTION_NO_FUZZER,
		"android.hardware.security.secretkeeper.ISecretkeeper/nonsecure":          EXCEPTION_NO_FUZZER,
		"android.hardware.security.secureclock.ISecureClock/default":              EXCEPTION_NO_FUZZER,
		"android.hardware.security.sharedsecret.ISharedSecret/default":            EXCEPTION_NO_FUZZER,
		"android.hardware.sensors.ISensors/default":                               EXCEPTION_NO_FUZZER,
		"android.hardware.soundtrigger3.ISoundTriggerHw/default":                  EXCEPTION_NO_FUZZER,
		"android.hardware.tetheroffload.IOffload/default":                         EXCEPTION_NO_FUZZER,
		"android.hardware.thermal.IThermal/default":                               EXCEPTION_NO_FUZZER,
		"android.hardware.threadnetwork.IThreadChip/chip0":                        []string{"android.hardware.threadnetwork-service.fuzzer"},
		"android.hardware.tv.hdmi.cec.IHdmiCec/default":                           EXCEPTION_NO_FUZZER,
		"android.hardware.tv.hdmi.connection.IHdmiConnection/default":             EXCEPTION_NO_FUZZER,
		"android.hardware.tv.hdmi.earc.IEArc/default":                             EXCEPTION_NO_FUZZER,
		"android.hardware.tv.input.ITvInput/default":                              EXCEPTION_NO_FUZZER,
		"android.hardware.tv.tuner.ITuner/default":                                EXCEPTION_NO_FUZZER,
		"android.hardware.usb.IUsb/default":                                       EXCEPTION_NO_FUZZER,
		"android.hardware.usb.gadget.IUsbGadget/default":                          EXCEPTION_NO_FUZZER,
		"android.hardware.uwb.IUwb/default":                                       EXCEPTION_NO_FUZZER,
		"android.hardware.vibrator.IVibrator/default":                             EXCEPTION_NO_FUZZER,
		"android.hardware.vibrator.IVibratorManager/default":                      []string{"android.hardware.vibrator-service.example_fuzzer"},
		"android.hardware.weaver.IWeaver/default":                                 EXCEPTION_NO_FUZZER,
		"android.hardware.wifi.IWifi/default":                                     EXCEPTION_NO_FUZZER,
		"android.hardware.wifi.hostapd.IHostapd/default":                          EXCEPTION_NO_FUZZER,
		"android.hardware.wifi.supplicant.ISupplicant/default":                    EXCEPTION_NO_FUZZER,
		"android.frameworks.cameraservice.service.ICameraService/default":         EXCEPTION_NO_FUZZER,
		"android.frameworks.location.altitude.IAltitudeService/default":           EXCEPTION_NO_FUZZER,
		"android.frameworks.sensorservice.ISensorManager/default":                 []string{"libsensorserviceaidl_fuzzer"},
		"android.frameworks.stats.IStats/default":                                 EXCEPTION_NO_FUZZER,
                "android.frameworks.vibrator.IVibratorControlService/default":             EXCEPTION_NO_FUZZER,
		"android.se.omapi.ISecureElementService/default":                          EXCEPTION_NO_FUZZER,
		"android.system.keystore2.IKeystoreService/default":                       EXCEPTION_NO_FUZZER,
		"android.system.net.netd.INetd/default":                                   []string{"netd_hw_service_fuzzer"},
		"android.system.suspend.ISystemSuspend/default":                           EXCEPTION_NO_FUZZER,
		"accessibility":      EXCEPTION_NO_FUZZER,
		"account":            EXCEPTION_NO_FUZZER,
		"activity":           EXCEPTION_NO_FUZZER,
		"activity_task":      EXCEPTION_NO_FUZZER,
		"adb":                EXCEPTION_NO_FUZZER,
		"adservices_manager": EXCEPTION_NO_FUZZER,
		"aidl_lazy_test_1":   EXCEPTION_NO_FUZZER,
		"aidl_lazy_test_2":   EXCEPTION_NO_FUZZER,
		"aidl_lazy_test_quit":   EXCEPTION_NO_FUZZER,
		"aidl_lazy_cb_test":  EXCEPTION_NO_FUZZER,
		"alarm":              EXCEPTION_NO_FUZZER,
		"android.hardware.automotive.evs.IEvsEnumerator/default":          EXCEPTION_NO_FUZZER,
		"android.os.UpdateEngineService":                                  []string{"update_engine_service_fuzzer"},
		"android.os.UpdateEngineStableService":                            []string{"update_engine_service_fuzzer"},
		"android.frameworks.automotive.display.ICarDisplayProxy/default":  EXCEPTION_NO_FUZZER,
		"android.security.apc":                                            EXCEPTION_NO_FUZZER,
		"android.security.authorization":                                  []string{"authorization_service_fuzzer"},
		"android.security.compat":                                         EXCEPTION_NO_FUZZER,
		"android.security.dice.IDiceMaintenance":                          EXCEPTION_NO_FUZZER,
		"android.security.dice.IDiceNode":                                 EXCEPTION_NO_FUZZER,
		"android.security.identity":                                       []string{"credstore_service_fuzzer"},
		"android.security.keystore":                                       EXCEPTION_NO_FUZZER,
		"android.security.legacykeystore":                                 EXCEPTION_NO_FUZZER,
		"android.security.maintenance":                                    EXCEPTION_NO_FUZZER,
		"android.security.metrics":                                        EXCEPTION_NO_FUZZER,
		"android.service.gatekeeper.IGateKeeperService":                   []string{"gatekeeperd_service_fuzzer"},
		"android.system.composd":                                          EXCEPTION_NO_FUZZER,
		// TODO(b/294158658): add fuzzer
		"android.hardware.security.keymint.IRemotelyProvisionedComponent/avf": EXCEPTION_NO_FUZZER,
		"android.system.virtualizationservice":                            EXCEPTION_NO_FUZZER,
		"android.system.virtualizationservice_internal.IVfioHandler":      EXCEPTION_NO_FUZZER,
		"ambient_context":                                                 EXCEPTION_NO_FUZZER,
		"app_binding":                                                     EXCEPTION_NO_FUZZER,
		"app_hibernation":                                                 EXCEPTION_NO_FUZZER,
		"app_integrity":                                                   EXCEPTION_NO_FUZZER,
		"app_prediction":                                                  EXCEPTION_NO_FUZZER,
		"app_search":                                                      EXCEPTION_NO_FUZZER,
		"apexservice":                                                     EXCEPTION_NO_FUZZER,
		"archive":                                                         EXCEPTION_NO_FUZZER,
		"attestation_verification":                                        EXCEPTION_NO_FUZZER,
		"blob_store":                                                      EXCEPTION_NO_FUZZER,
		"gsiservice":                                                      EXCEPTION_NO_FUZZER,
		"appops":                                                          EXCEPTION_NO_FUZZER,
		"appwidget":                                                       EXCEPTION_NO_FUZZER,
		"artd":                                                            EXCEPTION_NO_FUZZER,
		"assetatlas":                                                      EXCEPTION_NO_FUZZER,
		"attention":                                                       EXCEPTION_NO_FUZZER,
		"audio":                                                           EXCEPTION_NO_FUZZER,
		"auth":                                                            EXCEPTION_NO_FUZZER,
		"autofill":                                                        EXCEPTION_NO_FUZZER,
		"background_install_control":                                      EXCEPTION_NO_FUZZER,
		"backup":                                                          EXCEPTION_NO_FUZZER,
		"batteryproperties":                                               EXCEPTION_NO_FUZZER,
		"batterystats":                                                    EXCEPTION_NO_FUZZER,
		"battery":                                                         EXCEPTION_NO_FUZZER,
		"binder_calls_stats":                                              EXCEPTION_NO_FUZZER,
		"biometric":                                                       EXCEPTION_NO_FUZZER,
		"bluetooth_manager":                                               EXCEPTION_NO_FUZZER,
		"bluetooth":                                                       EXCEPTION_NO_FUZZER,
		"broadcastradio":                                                  EXCEPTION_NO_FUZZER,
		"bugreport":                                                       EXCEPTION_NO_FUZZER,
		"cacheinfo":                                                       EXCEPTION_NO_FUZZER,
		"carrier_config":                                                  EXCEPTION_NO_FUZZER,
		"clipboard":                                                       EXCEPTION_NO_FUZZER,
		"cloudsearch":                                                     EXCEPTION_NO_FUZZER,
		"cloudsearch_service":                                             EXCEPTION_NO_FUZZER,
		"com.android.net.IProxyService":                                   EXCEPTION_NO_FUZZER,
		"companiondevice":                                                 EXCEPTION_NO_FUZZER,
		"communal":                                                        EXCEPTION_NO_FUZZER,
		"platform_compat":                                                 EXCEPTION_NO_FUZZER,
		"platform_compat_native":                                          EXCEPTION_NO_FUZZER,
		"connectivity":                                                    EXCEPTION_NO_FUZZER,
		"connectivity_native":                                             EXCEPTION_NO_FUZZER,
		"connmetrics":                                                     EXCEPTION_NO_FUZZER,
		"consumer_ir":                                                     EXCEPTION_NO_FUZZER,
		"content":                                                         EXCEPTION_NO_FUZZER,
		"content_capture":                                                 EXCEPTION_NO_FUZZER,
		"content_suggestions":                                             EXCEPTION_NO_FUZZER,
		"contexthub":                                                      EXCEPTION_NO_FUZZER,
		"country_detector":                                                EXCEPTION_NO_FUZZER,
		"coverage":                                                        EXCEPTION_NO_FUZZER,
		"cpuinfo":                                                         EXCEPTION_NO_FUZZER,
		"cpu_monitor":                                                     EXCEPTION_NO_FUZZER,
		"credential":                                                      EXCEPTION_NO_FUZZER,
		"crossprofileapps":                                                EXCEPTION_NO_FUZZER,
		"dataloader_manager":                                              EXCEPTION_NO_FUZZER,
		"dbinfo":                                                          EXCEPTION_NO_FUZZER,
		"device_config":                                                   EXCEPTION_NO_FUZZER,
		"device_config_updatable":                                         EXCEPTION_NO_FUZZER,
		"device_policy":                                                   EXCEPTION_NO_FUZZER,
		"device_identifiers":                                              EXCEPTION_NO_FUZZER,
		"deviceidle":                                                      EXCEPTION_NO_FUZZER,
		"device_lock":                                                     EXCEPTION_NO_FUZZER,
		"device_state":                                                    EXCEPTION_NO_FUZZER,
		"devicestoragemonitor":                                            EXCEPTION_NO_FUZZER,
		"diskstats":                                                       EXCEPTION_NO_FUZZER,
		"display":                                                         EXCEPTION_NO_FUZZER,
		"dnsresolver":                                                     []string{"resolv_service_fuzzer"},
		"domain_verification":                                             EXCEPTION_NO_FUZZER,
		"color_display":                                                   EXCEPTION_NO_FUZZER,
		"netd_listener":                                                   EXCEPTION_NO_FUZZER,
		"network_watchlist":                                               EXCEPTION_NO_FUZZER,
		"DockObserver":                                                    EXCEPTION_NO_FUZZER,
		"dreams":                                                          EXCEPTION_NO_FUZZER,
		"drm.drmManager":                                                  []string{"drmserver_fuzzer"},
		"dropbox":                                                         EXCEPTION_NO_FUZZER,
		"dumpstate":                                                       EXCEPTION_NO_FUZZER,
		"dynamic_system":                                                  EXCEPTION_NO_FUZZER,
		"econtroller":                                                     EXCEPTION_NO_FUZZER,
		"emergency_affordance":                                            EXCEPTION_NO_FUZZER,
		"euicc_card_controller":                                           EXCEPTION_NO_FUZZER,
		"external_vibrator_service":                                       EXCEPTION_NO_FUZZER,
		"ethernet":                                                        EXCEPTION_NO_FUZZER,
		"face":                                                            EXCEPTION_NO_FUZZER,
		"file_integrity":                                                  EXCEPTION_NO_FUZZER,
		"fingerprint":                                                     EXCEPTION_NO_FUZZER,
		"feature_flags":                                                   EXCEPTION_NO_FUZZER,
		"font":                                                            EXCEPTION_NO_FUZZER,
		"android.hardware.fingerprint.IFingerprintDaemon": EXCEPTION_NO_FUZZER,
		"game":                         EXCEPTION_NO_FUZZER,
		"gfxinfo":                      EXCEPTION_NO_FUZZER,
		"gnss_time_update_service":     EXCEPTION_NO_FUZZER,
		"grammatical_inflection":       EXCEPTION_NO_FUZZER,
		"graphicsstats":                EXCEPTION_NO_FUZZER,
		"gpu":                          []string{"gpu_service_fuzzer"},
		"hardware":                     EXCEPTION_NO_FUZZER,
		"hardware_properties":          EXCEPTION_NO_FUZZER,
		"hdmi_control":                 EXCEPTION_NO_FUZZER,
		"healthconnect":                EXCEPTION_NO_FUZZER,
		"ions":                         EXCEPTION_NO_FUZZER,
		"idmap":                        EXCEPTION_NO_FUZZER,
		"incident":                     []string{"incidentd_service_fuzzer"},
		"incidentcompanion":            EXCEPTION_NO_FUZZER,
		"inputflinger":                 EXCEPTION_NO_FUZZER,
		"input_method":                 EXCEPTION_NO_FUZZER,
		"input":                        EXCEPTION_NO_FUZZER,
		"installd":                     []string{"installd_service_fuzzer"},
		"iphonesubinfo_msim":           EXCEPTION_NO_FUZZER,
		"iphonesubinfo2":               EXCEPTION_NO_FUZZER,
		"iphonesubinfo":                EXCEPTION_NO_FUZZER,
		"ims":                          EXCEPTION_NO_FUZZER,
		"imms":                         EXCEPTION_NO_FUZZER,
		"incremental":                  EXCEPTION_NO_FUZZER,
		"ipsec":                        EXCEPTION_NO_FUZZER,
		"ircsmessage":                  EXCEPTION_NO_FUZZER,
		"iris":                         EXCEPTION_NO_FUZZER,
		"isms_msim":                    EXCEPTION_NO_FUZZER,
		"isms2":                        EXCEPTION_NO_FUZZER,
		"isms":                         EXCEPTION_NO_FUZZER,
		"isub":                         EXCEPTION_NO_FUZZER,
		"jobscheduler":                 EXCEPTION_NO_FUZZER,
		"launcherapps":                 EXCEPTION_NO_FUZZER,
		"legacy_permission":            EXCEPTION_NO_FUZZER,
		"lights":                       EXCEPTION_NO_FUZZER,
		"locale":                       EXCEPTION_NO_FUZZER,
		"location":                     EXCEPTION_NO_FUZZER,
		"location_time_zone_manager":   EXCEPTION_NO_FUZZER,
		"lock_settings":                EXCEPTION_NO_FUZZER,
		"logcat":                       EXCEPTION_NO_FUZZER,
		"logd":                         EXCEPTION_NO_FUZZER,
		"looper_stats":                 EXCEPTION_NO_FUZZER,
		"lpdump_service":               EXCEPTION_NO_FUZZER,
		"mdns":                         EXCEPTION_NO_FUZZER,
		"media.aaudio":                 EXCEPTION_NO_FUZZER,
		"media.audio_flinger":          []string{"audioflinger_aidl_fuzzer"},
		"media.audio_policy":           []string{"audiopolicy_aidl_fuzzer"},
		"media.camera":                 []string{"camera_service_aidl_fuzzer"},
		"media.camera.proxy":           EXCEPTION_NO_FUZZER,
		"media.log":                    EXCEPTION_NO_FUZZER,
		"media.player":                 []string{"media_player_service_fuzzer"},
		"media.metrics":                []string{"mediametrics_aidl_fuzzer"},
		"media.extractor":              []string{"mediaextractor_service_fuzzer"},
		"media.transcoding":            EXCEPTION_NO_FUZZER,
		"media.resource_manager":       EXCEPTION_NO_FUZZER,
		"media.resource_observer":      EXCEPTION_NO_FUZZER,
		"media.sound_trigger_hw":       EXCEPTION_NO_FUZZER,
		"media.drm":                    EXCEPTION_NO_FUZZER,
		"media.tuner":                  EXCEPTION_NO_FUZZER,
		"media_communication":          EXCEPTION_NO_FUZZER,
		"media_metrics":                EXCEPTION_NO_FUZZER,
		"media_projection":             EXCEPTION_NO_FUZZER,
		"media_resource_monitor":       EXCEPTION_NO_FUZZER,
		"media_router":                 EXCEPTION_NO_FUZZER,
		"media_session":                EXCEPTION_NO_FUZZER,
		"meminfo":                      EXCEPTION_NO_FUZZER,
		"memtrack.proxy":               EXCEPTION_NO_FUZZER,
		"midi":                         EXCEPTION_NO_FUZZER,
		"mount":                        EXCEPTION_NO_FUZZER,
		"music_recognition":            EXCEPTION_NO_FUZZER,
		"nearby":                       EXCEPTION_NO_FUZZER,
		"netd":                         []string{"netd_native_service_fuzzer"},
		"netpolicy":                    EXCEPTION_NO_FUZZER,
		"netstats":                     EXCEPTION_NO_FUZZER,
		"network_stack":                EXCEPTION_NO_FUZZER,
		"network_management":           EXCEPTION_NO_FUZZER,
		"network_score":                EXCEPTION_NO_FUZZER,
		"network_time_update_service":  EXCEPTION_NO_FUZZER,
		"nfc":                          EXCEPTION_NO_FUZZER,
		"notification":                 EXCEPTION_NO_FUZZER,
		"oem_lock":                     EXCEPTION_NO_FUZZER,
		"ondevicepersonalization_system_service": EXCEPTION_NO_FUZZER,
		"otadexopt":                    EXCEPTION_NO_FUZZER,
		"ot_daemon":                    []string{"ot_daemon_service_fuzzer"},
		"overlay":                      EXCEPTION_NO_FUZZER,
		"pac_proxy":                    EXCEPTION_NO_FUZZER,
		"package":                      EXCEPTION_NO_FUZZER,
		"package_native":               EXCEPTION_NO_FUZZER,
		"people":                       EXCEPTION_NO_FUZZER,
		"performance_hint":             EXCEPTION_NO_FUZZER,
		"permission":                   EXCEPTION_NO_FUZZER,
		"permissionmgr":                EXCEPTION_NO_FUZZER,
		"permission_checker":           EXCEPTION_NO_FUZZER,
		"persistent_data_block":        EXCEPTION_NO_FUZZER,
		"phone_msim":                   EXCEPTION_NO_FUZZER,
		"phone1":                       EXCEPTION_NO_FUZZER,
		"phone2":                       EXCEPTION_NO_FUZZER,
		"phone":                        EXCEPTION_NO_FUZZER,
		"pinner":                       EXCEPTION_NO_FUZZER,
		"powerstats":                   EXCEPTION_NO_FUZZER,
		"power":                        EXCEPTION_NO_FUZZER,
		"print":                        EXCEPTION_NO_FUZZER,
		"processinfo":                  EXCEPTION_NO_FUZZER,
		"procstats":                    EXCEPTION_NO_FUZZER,
		"profcollectd":                 EXCEPTION_NO_FUZZER,
		"radio.phonesubinfo":           EXCEPTION_NO_FUZZER,
		"radio.phone":                  EXCEPTION_NO_FUZZER,
		"radio.sms":                    EXCEPTION_NO_FUZZER,
		"rcs":                          EXCEPTION_NO_FUZZER,
		"reboot_readiness":             EXCEPTION_NO_FUZZER,
		"recovery":                     EXCEPTION_NO_FUZZER,
		"remote_auth":                  EXCEPTION_NO_FUZZER,
		"remote_provisioning":          EXCEPTION_NO_FUZZER,
		"resolver":                     EXCEPTION_NO_FUZZER,
		"resources":                    EXCEPTION_NO_FUZZER,
		"restrictions":                 EXCEPTION_NO_FUZZER,
		"rkpd.registrar":               EXCEPTION_NO_FUZZER,
		"rkpd.refresh":                 EXCEPTION_NO_FUZZER,
		"role":                         EXCEPTION_NO_FUZZER,
		"rollback":                     EXCEPTION_NO_FUZZER,
		"rttmanager":                   EXCEPTION_NO_FUZZER,
		"runtime":                      EXCEPTION_NO_FUZZER,
		"safety_center":                EXCEPTION_NO_FUZZER,
		"samplingprofiler":             EXCEPTION_NO_FUZZER,
		"scheduling_policy":            EXCEPTION_NO_FUZZER,
		"search":                       EXCEPTION_NO_FUZZER,
		"search_ui":                    EXCEPTION_NO_FUZZER,
		"secure_element":               EXCEPTION_NO_FUZZER,
		"security_state":               EXCEPTION_NO_FUZZER,
		"sec_key_att_app_id_provider":  EXCEPTION_NO_FUZZER,
		"selection_toolbar":            EXCEPTION_NO_FUZZER,
		"sensorservice":                EXCEPTION_NO_FUZZER,
		"sensor_privacy":               EXCEPTION_NO_FUZZER,
		"serial":                       EXCEPTION_NO_FUZZER,
		"servicediscovery":             EXCEPTION_NO_FUZZER,
		"manager":                      []string{"servicemanager_fuzzer"},
		"settings":                     EXCEPTION_NO_FUZZER,
		"shortcut":                     EXCEPTION_NO_FUZZER,
		"simphonebook_msim":            EXCEPTION_NO_FUZZER,
		"simphonebook2":                EXCEPTION_NO_FUZZER,
		"simphonebook":                 EXCEPTION_NO_FUZZER,
		"sip":                          EXCEPTION_NO_FUZZER,
		"slice":                        EXCEPTION_NO_FUZZER,
		"smartspace":                   EXCEPTION_NO_FUZZER,
		"speech_recognition":           EXCEPTION_NO_FUZZER,
		"stats":                        EXCEPTION_NO_FUZZER,
		"statsbootstrap":               EXCEPTION_NO_FUZZER,
		"statscompanion":               EXCEPTION_NO_FUZZER,
		"statsmanager":                 EXCEPTION_NO_FUZZER,
		"soundtrigger":                 EXCEPTION_NO_FUZZER,
		"soundtrigger_middleware":      EXCEPTION_NO_FUZZER,
		"statusbar":                    EXCEPTION_NO_FUZZER,
		"storaged":                     []string{"storaged_service_fuzzer"},
		"storaged_pri":                 []string{"storaged_private_service_fuzzer"},
		"storagestats":                 EXCEPTION_NO_FUZZER,
		"sdk_sandbox":                  EXCEPTION_NO_FUZZER,
		"SurfaceFlinger":               EXCEPTION_NO_FUZZER,
		"SurfaceFlingerAIDL":           EXCEPTION_NO_FUZZER,
		"suspend_control":              []string{"suspend_service_fuzzer"},
		"suspend_control_internal":     []string{"suspend_service_internal_fuzzer"},
		"system_config":                EXCEPTION_NO_FUZZER,
		"system_server_dumper":         EXCEPTION_NO_FUZZER,
		"system_update":                EXCEPTION_NO_FUZZER,
		"tare":                         EXCEPTION_NO_FUZZER,
		"task":                         EXCEPTION_NO_FUZZER,
		"telecom":                      EXCEPTION_NO_FUZZER,
		"telephony.registry":           EXCEPTION_NO_FUZZER,
		"telephony_ims":                EXCEPTION_NO_FUZZER,
		"testharness":                  EXCEPTION_NO_FUZZER,
		"tethering":                    EXCEPTION_NO_FUZZER,
		"textclassification":           EXCEPTION_NO_FUZZER,
		"textservices":                 EXCEPTION_NO_FUZZER,
		"texttospeech":                 EXCEPTION_NO_FUZZER,
		"thread_network":               EXCEPTION_NO_FUZZER,
		"time_detector":                EXCEPTION_NO_FUZZER,
		"time_zone_detector":           EXCEPTION_NO_FUZZER,
		"thermalservice":               EXCEPTION_NO_FUZZER,
		"tracing.proxy":                EXCEPTION_NO_FUZZER,
		"translation":                  EXCEPTION_NO_FUZZER,
		"transparency":                 EXCEPTION_NO_FUZZER,
		"trust":                        EXCEPTION_NO_FUZZER,
		"tv_ad":                        EXCEPTION_NO_FUZZER,
		"tv_interactive_app":           EXCEPTION_NO_FUZZER,
		"tv_input":                     EXCEPTION_NO_FUZZER,
		"tv_tuner_resource_mgr":        EXCEPTION_NO_FUZZER,
		"uce":                          EXCEPTION_NO_FUZZER,
		"uimode":                       EXCEPTION_NO_FUZZER,
		"updatelock":                   EXCEPTION_NO_FUZZER,
		"uri_grants":                   EXCEPTION_NO_FUZZER,
		"usagestats":                   EXCEPTION_NO_FUZZER,
		"usb":                          EXCEPTION_NO_FUZZER,
		"user":                         EXCEPTION_NO_FUZZER,
		"uwb":                          EXCEPTION_NO_FUZZER,
		"vcn_management":               EXCEPTION_NO_FUZZER,
		"vibrator":                     EXCEPTION_NO_FUZZER,
		"vibrator_manager":             EXCEPTION_NO_FUZZER,
		"virtualdevice":                EXCEPTION_NO_FUZZER,
		"virtualdevice_native":         EXCEPTION_NO_FUZZER,
		"virtual_camera":               EXCEPTION_NO_FUZZER,
		"virtual_touchpad":             EXCEPTION_NO_FUZZER,
		"voiceinteraction":             EXCEPTION_NO_FUZZER,
		"vold":                         []string{"vold_native_service_fuzzer"},
		"vpn_management":               EXCEPTION_NO_FUZZER,
		"vrmanager":                    EXCEPTION_NO_FUZZER,
		"wallpaper":                    EXCEPTION_NO_FUZZER,
		"wallpaper_effects_generation": EXCEPTION_NO_FUZZER,
		"wearable_sensing":             EXCEPTION_NO_FUZZER,
		"webviewupdate":                EXCEPTION_NO_FUZZER,
		"wifip2p":                      EXCEPTION_NO_FUZZER,
		"wifiscanner":                  EXCEPTION_NO_FUZZER,
		"wifi":                         EXCEPTION_NO_FUZZER,
		"wifinl80211":                  []string{"wificond_service_fuzzer"},
		"wifiaware":                    EXCEPTION_NO_FUZZER,
		"wifirtt":                      EXCEPTION_NO_FUZZER,
		"window":                       EXCEPTION_NO_FUZZER,
		"*":                            EXCEPTION_NO_FUZZER,
	}
)
