/*
 * Copyright 2018, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */
package io.enmasse.systemtest.bases.infra;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.Arrays;
import java.util.List;
import java.util.Optional;
import java.util.function.Supplier;
import java.util.stream.Collectors;

import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.opentest4j.AssertionFailedError;
import org.slf4j.Logger;

import io.enmasse.systemtest.AddressSpace;
import io.enmasse.systemtest.CustomLogger;
import io.enmasse.systemtest.PlansProvider;
import io.enmasse.systemtest.TestUtils;
import io.enmasse.systemtest.TimeoutBudget;
import io.enmasse.systemtest.ability.ITestBase;
import io.enmasse.systemtest.bases.TestBase;
import io.enmasse.systemtest.resources.AddressPlanDefinition;
import io.enmasse.systemtest.resources.InfraConfigDefinition;
import io.enmasse.systemtest.resources.InfraSpecComponent;
import io.fabric8.kubernetes.api.model.Container;
import io.fabric8.kubernetes.api.model.PersistentVolumeClaim;
import io.fabric8.kubernetes.api.model.Pod;
import io.fabric8.kubernetes.api.model.ResourceRequirements;
import io.fabric8.kubernetes.api.model.storage.StorageClass;

public abstract class InfraTestBase extends TestBase implements ITestBase{

    protected static Logger log = CustomLogger.getLogger();

    private static final List<String> resizingStorageProvisioners = Arrays.asList("kubernetes.io/aws-ebs", "kubernetes.io/gce-pd",
            "kubernetes.io/azure-file", "kubernetes.io/azure-disk", "kubernetes.io/glusterfs", "kubernetes.io/cinder",
            "kubernetes.io/portworx-volume", "kubernetes.io/rbd");

    protected static final PlansProvider plansProvider = new PlansProvider(kubernetes);

    protected InfraConfigDefinition testInfra;
    protected AddressPlanDefinition exampleAddressPlan;
    protected AddressSpace exampleAddressSpace;

    @BeforeEach
    void setUp() throws Exception {
        plansProvider.setUp();
    }

    @AfterEach
    void tearDown() throws Exception {
        plansProvider.tearDown();
    }

    protected void assertBroker(String brokerMemory, Optional<String> brokerStorage) {
        log.info("Checking broker infra");
        List<Pod> brokerPods = TestUtils.listBrokerPods(kubernetes, exampleAddressSpace);
        assertEquals(1, brokerPods.size());

        Pod broker = brokerPods.stream().findFirst().get();
        String actualBrokerMemory = broker.getSpec().getContainers().stream()
                .filter(container->container.getName().equals("broker")).findFirst()
                .map(Container::getResources)
                .map(ResourceRequirements::getLimits)
                .get().get("memory").getAmount();
        assertEquals(brokerMemory, actualBrokerMemory, "Broker memory limit incorrect");

        if(brokerStorage.isPresent()) {
            PersistentVolumeClaim brokerVolumeClaim = getBrokerPVCData(broker);
            assertEquals(brokerStorage.get(), brokerVolumeClaim.getSpec().getResources().getRequests().get("storage").getAmount(),
                    "Broker data storage request incorrect");
        }
    }

    protected void assertAdminConsole(String adminMemory) {
        log.info("Checking admin console infra");
        List<Pod> adminPods = TestUtils.listAdminConsolePods(kubernetes, exampleAddressSpace);
        assertEquals(1, adminPods.size());

        List<ResourceRequirements> adminResources = adminPods.stream().findFirst().get().getSpec().getContainers()
                .stream().map(Container::getResources).collect(Collectors.toList());

        for (ResourceRequirements requirements : adminResources) {
            assertEquals(adminMemory, requirements.getLimits().get("memory").getAmount(),
                    "Admin console memory limit incorrect");
            assertEquals(adminMemory, requirements.getRequests().get("memory").getAmount(),
                    "Admin console memory requests incorrect");
        }
    }

    protected void waitUntilInfraReady(Supplier<Boolean> assertCall, TimeoutBudget timeout) throws InterruptedException {
        log.info("Start waiting for infra ready");
        AssertionFailedError lastException = null;
        while (!timeout.timeoutExpired()) {
            try {
                assertCall.get();
                log.info("assert infra ready succeed");
                return;
            }catch(AssertionFailedError e) {
                lastException = e;
            }
            log.debug("next iteration, remaining time: {}", timeout.timeLeft());
            Thread.sleep(5000);
        }
        log.error("Timeout assert infra expired");
        if(lastException!=null) {
            throw lastException;
        }
    }

    protected PersistentVolumeClaim getBrokerPVCData(Pod broker) {
        String brokerVolumeClaimName = broker.getSpec().getVolumes().stream()
                .filter(volume->volume.getName().equals("data"))
                .findFirst().get()
                .getPersistentVolumeClaim().getClaimName();
        PersistentVolumeClaim brokerVolumeClaim = TestUtils.listPersistentVolumeClaims(kubernetes, exampleAddressSpace).stream()
                .filter(pvc->pvc.getMetadata().getName().equals(brokerVolumeClaimName))
                .findFirst().get();
        return brokerVolumeClaim;
    }

    protected InfraSpecComponent getInfraComponent(InfraConfigDefinition infra, String type) {
        return infra.getAddressResources().stream().filter(isc->isc.getType().equals(type)).findFirst().get();
    }

    protected Boolean volumeResizingSupported() {
        List<Pod> brokerPods = TestUtils.listBrokerPods(kubernetes, exampleAddressSpace);
        assertEquals(1, brokerPods.size());
        Pod broker = brokerPods.stream().findFirst().get();
        PersistentVolumeClaim brokerVolumeClaim = getBrokerPVCData(broker);
        String brokerStorageClassName = brokerVolumeClaim.getSpec().getStorageClassName();
        if(brokerStorageClassName!=null) {
            StorageClass brokerStorageClass = kubernetes.getStorageClass(brokerStorageClassName);
            if(resizingStorageProvisioners.contains(brokerStorageClass.getProvisioner())) {
                if(brokerStorageClass.getAllowVolumeExpansion()!=null && brokerStorageClass.getAllowVolumeExpansion()) {
                    log.info("Testing broker volume resize because of {}:{}", brokerStorageClassName, brokerStorageClass.getProvisioner());
                    return true;
                }else {
                    log.info("Skipping broker volume resize due to allowVolumeExpansion in StorageClass {} disabled", brokerStorageClassName);
                }
            }else {
                log.info("Skipping broker volume resize due to provisioner: {}", brokerStorageClass.getProvisioner());
            }
        }else {
            log.info("Skipping broker volume resize due to missing StorageClass name in PVC {}", brokerVolumeClaim.getMetadata().getName());
        }
        return false;
    }

}
